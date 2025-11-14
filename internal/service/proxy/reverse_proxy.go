package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"uniauth-gateway/internal/service/uniGf"
)

func ReverseProxy(r *ghttp.Request) {
	ctx := r.Context()

	// 规则一，查看请求头请求的服务
	service := r.Header.Get("X-Service")
	if service == "" {
		// 规则二，路由匹配，常用于前端SPA的返回
		pathList := strings.Split(r.URL.Path, "/")
		if len(pathList) < 2 || pathList[1] == "" {
			// 规则三，判断配置是否允许找不到服务时自动重定向到 /chat/
			if g.Cfg().MustGet(ctx, "server.allowAutoRedirectToChat", gvar.New(false)).Bool() {
				r.Response.RedirectTo("/chat/")
				return
			}
			r.Response.WriteStatusExit(http.StatusBadRequest, "未获取到服务名称")
		}
		service = pathList[1]
	}

	// 检查微服务权限
	if err := uniGf.CheckPermission(ctx, r.Session.MustGet("user_id").String(), service, "entry"); err != nil {
		r.Response.WriteStatusExit(http.StatusForbidden, err)
		return
	}

	// 获取服务对应的代理地址
	proxyHostMap := g.Cfg("proxy_host_map").MustData(ctx)
	proxyHostVar, ok := proxyHostMap[service]
	if !ok {
		g.Log().Errorf(r.Context(), "未找到服务[%s]对应的代理地址", service)
		r.Response.WriteStatus(http.StatusBadRequest, fmt.Sprintf("未找到服务[%s]对应的代理地址", service))
		return
	}
	proxyHost := proxyHostVar.(string)

	// 创建反向代理
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			// 解析目标地址
			target, err := url.Parse(proxyHost)
			if err != nil {
				g.Log().Errorf(r.Context(), "Proxy URL %s 解析失败：%v", proxyHost, err)
				return
			}

			// 重写请求的 URL、Host 和 请求头
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			// 如果路径不以斜杠结尾，自动补上（避免 Nginx 301 补斜杠）
			// 只有同时满足两个条件才补斜杠：
			// 路径不以 / 结尾（!strings.HasSuffix(forwardPath, "/")）
			// 最后一段不包含 .（!strings.Contains(..., ".")）
			forwardPath := r.URL.Path
			if !strings.HasSuffix(forwardPath, "/") && !strings.Contains(forwardPath[strings.LastIndex(forwardPath, "/")+1:], ".") {
				forwardPath += "/"
			}
			req.URL.Path = forwardPath
			req.Host = r.Host
			if prior := req.Header.Get("X-Forwarded-For"); prior != "" {
				req.Header.Set("X-Forwarded-For", prior+", "+r.RemoteAddr)
			} else {
				req.Header.Set("X-Forwarded-For", r.RemoteAddr)
			}
			req.Header.Set("X-Forwarded-Host", r.Host)
			if r.TLS != nil {
				req.Header.Set("X-Forwarded-Proto", "https")
			} else {
				req.Header.Set("X-Forwarded-Proto", "http")
			}

			// 鉴权信息嵌入
			req.Header.Set("X-User-ID", r.Session.MustGet("user_id").String())

			g.Log().Infof(r.Context(), `[网关]: %s -> [%s]: %s://%s%s`, r.GetUrl(), service, req.URL.Scheme, req.URL.Host, req.URL.Path)
		},
		ModifyResponse: func(resp *http.Response) error {
			// 记录响应状态和 Location
			if resp.StatusCode >= 300 && resp.StatusCode < 400 {
				location := resp.Header.Get("Location")
				g.Log().Infof(r.Context(), `[%s 上游返回重定向] Status=%d, Location=%s。请求URL: %s`, service, resp.StatusCode, location, r.URL.String())
			}
			return nil
		},
		ErrorHandler: func(writer http.ResponseWriter, request *http.Request, e error) {
			g.Log().Errorf(r.Context(), "proxy 失败: %v", e)
			writer.WriteHeader(http.StatusBadGateway)
		},
	}
	proxy.ServeHTTP(r.Response.Writer, r.Request)
}
