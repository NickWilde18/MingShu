package middlewares

import (
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ReturnSessionTTL(r *ghttp.Request) {
	sessionId, err := r.Session.Id()
	userId, errGetUser := r.Session.Get("user_id")
	if err != nil {
		g.Log().Error(r.Context(), gerror.Wrapf(err, "获取会话 ID 失败。会话存储的 UPN 为：%s", userId.String()))
		r.Middleware.Next()
		return
	}
	if errGetUser != nil {
		g.Log().Error(r.Context(), gerror.Wrap(errGetUser, "会话存储的 UPN 获取失败。"))
		// 这里不用返回，给个默认值
		userId = gvar.New("未找到")
	}

	if TTL, errTTL := g.Redis().Do(r.Context(), "TTL", sessionId); errTTL != nil {
		g.Log().Error(r.Context(), gerror.Wrapf(errTTL, "访问 Redis 查询会话 TTL 失败。会话 ID 为：%s。用户 UPN 为：%s", sessionId, userId.String()))
		// 失败也不影响后续处理
	} else {
		// 在业务 Handler 写响应之前设置好 Header
		r.Response.Header().Set("X-Session-Expires-In", TTL.String())
	}

	// 继续后续中间件和业务 Handler
	r.Middleware.Next()
}
