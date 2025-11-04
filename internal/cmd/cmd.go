package cmd

import (
	"context"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"

	"uniauth-gateway/internal/middlewares"
	"uniauth-gateway/internal/service/auth"
	"uniauth-gateway/internal/service/proxy"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// ghttp server 各种特性开启与配置
			// === 通用配置 ===
			// 标记是否是本地环境。
			local := g.Cfg().MustGet(ctx, "server.local", false).Bool()
			// 日志等级设置
			if local {
				if err := g.Log().SetLevelStr("all"); err != nil {
					g.Log().Errorf(ctx, "设置日志等级 all 错误：%v", err)
				}
			} else {
				if err := g.Log().SetLevelStr("INFO"); err != nil {
					g.Log().Errorf(ctx, "设置日志等级 INFO 错误：%v", err)
				}
			}
			// 设置 HTTP 服务器
			s.SetPort(g.Cfg().MustGet(ctx, "server.httpPort").Int())
			// 设置 Session Age
			s.SetSessionMaxAge(12 * time.Hour)
			// === 集群内外部启动配置 ===
			if !local {
				// 集群外部启动
				// 启动 HTTPS 服务器
				s.EnableHTTPS(
					"/app/certs/tls.crt",
					"/app/certs/tls.key",
				)
				s.SetHTTPSPort(g.Cfg().MustGet(ctx, "server.httpsPort").Int())
				// 使用内存作为 Session 存储位置
				s.SetSessionStorage(gsession.NewStorageRedis(g.Redis()))
			} else {
				// 集群内部启动
				// 使用 Redis 作为 Session 存储位置
				s.SetSessionStorage(gsession.NewStorageMemory())
			}

			// 全局中间件 - 错误处理
			s.Use(middlewares.ErrorHandler)

			// 不需要登录验证的路由组
			s.Group("/auth", func(group *ghttp.RouterGroup) {
				group.GET("/login/*", auth.Login)
				group.GET("/logout", auth.Logout)
				group.GET("/callback", auth.Callback)
			})

			// 需要登录验证的路由组
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(middlewares.VerifyLoginStatus)
				group.ALL("/*", proxy.ReverseProxy)
			})

			s.Run()
			return nil
		},
	}
)
