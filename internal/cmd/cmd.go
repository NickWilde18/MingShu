package cmd

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	svc "MingShu/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.BindHandler("/", func(r *ghttp.Request) {
				// 规则一，查看请求头请求的服务
				service := r.Header.Get("X-Service")
				if service != "" {
					proxyHost, _ := svc.GetProxyHost(service)
					g.Log().Infof(r.Context(), `proxy:"%s" -> backend:"%s"`, r.Request.URL.Host, proxyHost)
					r.MakeBodyRepeatableRead(false)

					// 创建反向代理
					proxy := &httputil.ReverseProxy{
						Director: func(req *http.Request) {
							// 解析目标地址
							target, err := url.Parse(proxyHost)
							if err != nil {
								g.Log().Errorf(r.Context(), "Proxy URL %s 解析失败：%v", proxyHost, err)
								return
							}

							// 重写请求的 URL 和 Host
							req.URL.Scheme = target.Scheme
							req.URL.Host = target.Host
							req.Host = target.Host // 设置 Host 头
							g.Log().Infof(req.Context(), `proxy URL: %s`, req.URL.String())
							// 可选：添加自定义 Header
							req.Header.Set("X-Forwarded-For", req.RemoteAddr)
						},
						ErrorHandler: func(writer http.ResponseWriter, request *http.Request, e error) {
							g.Log().Errorf(r.Context(), "proxy 失败: %v", e)
							writer.WriteHeader(http.StatusBadGateway)
						},
					}
					proxy.ServeHTTP(r.Response.Writer, r.Request)
				}
				// 规则二，路由匹配
			})
			s.SetPort(10200)
			s.Run()
			return nil
		},
	}
)
