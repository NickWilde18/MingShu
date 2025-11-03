package cmd

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gogf/gf/contrib/config/kubecm/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 初始化，加载 ConfigMap
			adapter, err := kubecm.New(gctx.GetInitCtx(), kubecm.Config{
				ConfigMap: "dev-mingshu-gateway-config", // Name of the ConfigMap to use
				DataItem:  "proxy_host_map",             // Key in the ConfigMap data field
			})
			if err != nil {
				g.Log().Errorf(ctx, "从 Kuebernetes ConfigMap 初始化配置中心失败: %v", err)
				g.Log().Info(ctx, "从本地配置文件初始化配置中心")
				// return
			} else {
				g.Cfg("mingshu-config").SetAdapter(adapter)
				g.Log().Info(ctx, "从 Kubernetes ConfigMap 初始化配置中心")
			}

			// 服务器
			s := g.Server()
			s.BindHandler("/*", func(r *ghttp.Request) {
				g.Log().Infof(r.Context(), "请求URL: %s", r.GetUrl())
				// 规则一，查看请求头请求的服务
				service := r.Header.Get("X-Service")
				if service == "" {
					// 规则二，路由匹配，常用于前端SPA的返回
					pathList := strings.Split(r.URL.Path, "/")
					if len(pathList) < 2 || pathList[1] == "" {
						r.Response.WriteStatus(http.StatusBadRequest, "未获取到服务名称")
						return
					}
					service = pathList[1]
					g.Log().Infof(ctx, "从路径中获取服务名称: %s", service)
				}

				// 获取服务对应的代理地址
				proxyHostMap := g.Cfg("mingshu-config").MustData(ctx)
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
						req.URL.Path = r.URL.Path
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

						g.Log().Infof(r.Context(), `[Gateway]: %s -> [%s]: %s://%s%s`, r.GetUrl(), service, req.URL.Scheme, req.URL.Host, req.URL.Path)
					},
					ModifyResponse: func(resp *http.Response) error {
						// 记录响应状态和 Location
						if resp.StatusCode >= 300 && resp.StatusCode < 400 {
							location := resp.Header.Get("Location")
							g.Log().Infof(r.Context(), `[上游返回重定向] Status=%d, Location=%s`, resp.StatusCode, location)
						}
						return nil
					},
					ErrorHandler: func(writer http.ResponseWriter, request *http.Request, e error) {
						g.Log().Errorf(r.Context(), "proxy 失败: %v", e)
						writer.WriteHeader(http.StatusBadGateway)
					},
				}
				proxy.ServeHTTP(r.Response.Writer, r.Request)
			})

			// s.EnableHTTPS(
			// 	"/app/certs/tls.crt",
			// 	"/app/certs/tls.key",
			// )
			s.SetPort(8080)
			// s.SetHTTPSPort(8081)
			s.Run()
			return nil
		},
	}
)
