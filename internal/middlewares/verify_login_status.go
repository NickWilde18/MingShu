package middlewares

import (
	"uniauth-gateway/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// VerifyLoginStatus 登录状态验证中间件
func VerifyLoginStatus(r *ghttp.Request) {
	local := g.Cfg().MustGet(r.Context(), "server.local", false).Bool()
	if local {
		r.Session.Set("user_id", "mszure")
		r.Middleware.Next()
	}

	// 验证登录状态
	userID, err := r.Session.Get("user_id")
	if err != nil || userID.IsNil() {
		// 使用自定义错误页面，带自动跳转功能
		RenderError(r, ErrorInfo{
			ErrorCode: consts.ErrCodeUnauthorized,
			CustomMsg: `您尚未登录，系统将在3秒后自动跳转到登录页面...
You are not logged in, redirecting to login page in 3 seconds...`,
			CustomJS:  `setTimeout(function(){window.location.href='/auth/login';},3000);`,
		})
		return
	}
	r.Middleware.Next()
}
