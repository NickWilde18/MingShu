package auth

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/errors/gerror"
	
	m "uniauth-gateway/internal/middlewares"
	"uniauth-gateway/internal/consts"
)

func Logout(r *ghttp.Request) {
	if err := r.Session.RemoveAll(); err != nil {
		g.Log().Error(r.Context(), gerror.Wrap(err, "清除 Session 失败"))
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeLogoutFailed,
			Detail:    err.Error(),
		})
	}
	r.Response.RedirectTo(g.Cfg().MustGetWithEnv(r.Context(), "sso.logoutUrl").String())
}
