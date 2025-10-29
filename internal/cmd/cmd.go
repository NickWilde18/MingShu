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
			s.BindHandler("/*", func(r *ghttp.Request) {
				// 规则一，查看请求头请求的服务
				service := r.Header.Get("X-Service")
				if service != "" {
					proxyHost, err := svc.GetProxyHost(service)
					if err != nil {
						r.Response.WriteStatus(http.StatusBadRequest, err.Error())
						return
					}
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

							// 重写请求的 URL、Host 和 请求头
							req.URL.Scheme = target.Scheme
							req.URL.Host = target.Host
							req.URL.Path = r.URL.Path
							req.Host = target.Host
							req.Header.Set("X-Forwarded-For", req.RemoteAddr)

							g.Log().Infof(r.Context(), `[Gateway]: %s -> [%s]: %s://%s%s`, r.GetUrl(), service, req.URL.Scheme, req.URL.Host, req.URL.Path)
						},
						ErrorHandler: func(writer http.ResponseWriter, request *http.Request, e error) {
							g.Log().Errorf(r.Context(), "proxy 失败: %v", e)
							writer.WriteHeader(http.StatusBadGateway)
						},
					}
					proxy.ServeHTTP(r.Response.Writer, r.Request)
				} else {
					// 规则二，路由匹配

				}
			})
			s.SetPort(10200)
			s.Run()
			return nil
		},
	}
)
