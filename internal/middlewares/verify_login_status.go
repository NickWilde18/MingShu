package middlewares

import (
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

// VerifyLoginStatus 登录状态验证中间件
func VerifyLoginStatus(r *ghttp.Request) {
	// 验证登录状态
	userID, err := r.Session.Get("user_id")
	if err != nil || userID.IsNil() {
		// 使用自定义 JS 实现自动跳转到登录页面
		RenderError(r, ErrorInfo{
			Code:    http.StatusUnauthorized,
			Message: "您尚未登录，<a href='/auth/login'>点击这里登录</a>，或等待自动跳转...",
			CustomJS: `
// 3秒后自动跳转到登录页面
var countdown = 3;
var timer = setInterval(function() {
    countdown--;
    if (countdown <= 0) {
        clearInterval(timer);
        window.location.href = '/auth/login';
    }
}, 1000);
`,
		})
		return
	}
	r.Middleware.Next()
}
