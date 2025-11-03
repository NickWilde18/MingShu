package middlewares

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

func VerifyLoginStatus(r *ghttp.Request) {
	userID, err := r.Session.Get("user_id")
	if err != nil || userID.IsNil() {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "未登录")
	}
	r.Middleware.Next()
}