package uniGf

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	"uniauth-gateway/internal/consts"
	m "uniauth-gateway/internal/middlewares"
)

// CheckPermission 检查用户权限
func CheckPermission(ctx context.Context, sub string, obj string, act string) {
	response := client.PostVar(ctx, "/auth/check", g.Map{
		"sub": sub,
		"obj": obj,
		"act": act,
	})
	if response.IsNil() || response.IsEmpty() {
		m.RenderError(g.RequestFromCtx(ctx), m.ErrorInfo{
			ErrorCode: consts.ErrCodeBadGateway,
			Detail: `权限检查：统一鉴权内部请求没有响应或返回内容为空。
The permission check: the internal request to the unified authorization does not respond or the returned content is empty.`,
		})
	}
	content := response.Map()
	if !content["success"].(bool) {
		m.RenderError(g.RequestFromCtx(ctx), m.ErrorInfo{
			ErrorCode: consts.ErrCodeBadGateway,
			Detail: "权限检查失败：" + content["message"].(string) + `
The permission check failed: ` + content["message"].(string),
		})
	}
	if !content["data"].(g.Map)["allow"].(bool) {
		m.RenderError(g.RequestFromCtx(ctx), m.ErrorInfo{
			ErrorCode: consts.ErrCodeForbidden,
			Detail: fmt.Sprintf(`你没有权限使用 [%s] %s。请联系 ITSO。
You do not have permission to use [%s] %s. Please contact ITSO.`, obj, act, obj, act),
		})
	}
}
