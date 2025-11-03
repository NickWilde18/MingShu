package cmd

import (
	"context"
	"time"

	"github.com/gogf/gf/contrib/config/kubecm/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gsession"

	"MingShu/internal/middlewares"
	"MingShu/internal/service/proxy"
	"MingShu/internal/service/auth"
)

// 标记是否是本地环境。本地环境不启动 HTTPS 服务器。
var LOCAL bool = false

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			// 初始化
			// 配置配置文件来源，同时确定程序在集群内启动还是外部启动
			adapter, err := kubecm.New(gctx.GetInitCtx(), kubecm.Config{
				ConfigMap: "dev-mingshu-gateway-config",
				DataItem:  "proxy_host_map",
			})
			if err != nil {
				g.Log().Debugf(ctx, "从 Kuebernetes ConfigMap 初始化配置中心失败: %v", err)
				g.Log().Info(ctx, "从 本地配置文件 初始化配置中心")
				LOCAL = true
			} else {
				g.Cfg("mingshu-config").SetAdapter(adapter)
				g.Log().Info(ctx, "从 Kubernetes ConfigMap 初始化配置中心")
			}
			// ghttp server 各种特性开启与配置
			// === 通用配置 ===
			// 设置 HTTP 服务器
			s.SetPort(g.Cfg().MustGet(ctx, "server.httpPort").Int())
			// 设置 Session Age
			s.SetSessionMaxAge(12 * time.Hour)
			// === 集群内外部启动配置 ===
			if !LOCAL {
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

			// 中间件绑定
			s.Use(middlewares.ErrorHandler)
			s.Use(middlewares.VerifyLoginStatus)

			// Handler 绑定
			s.BindHandler("/auth/login", auth.Login)
			s.BindHandler("/auth/logout", auth.Logout)
			s.BindHandler("/auth/callback", auth.Callback)
			s.BindHandler("/*", proxy.ReverseProxy)

			s.Run()
			return nil
		},
	}
)
