package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func Logout(r *ghttp.Request) {
	if err := r.Session.RemoveAll(); err != nil {
		g.Log().Errorf(r.Context(), "清除 Session 失败: %v", err.Error())
	}
	r.Response.RedirectTo(g.Cfg().MustGetWithEnv(r.Context(), "sso.logout_url").String())
}
