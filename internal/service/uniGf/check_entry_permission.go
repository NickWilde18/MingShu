package uniGf

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// CheckEntryPermission 检查用户入口权限
func CheckEntryPermission(ctx context.Context, upn string) error {
	response := client.PostVar(ctx, "/auth/check", g.Map{
		"sub": upn,
		"obj": "platform",
		"act": "entry",
	})
	if response.IsNil() || response.IsEmpty() {
		return gerror.New("入口权限检查：内部请求没有响应或返回内容为空")
	}
	content := response.Map()
	if !content["success"].(bool) {
		return gerror.New("权限检查失败：" + content["message"].(string))
	}
	if !content["data"].(g.Map)["allow"].(bool) {
		return gerror.New("你没有权限进入平台。请联系管理员。")
	}
	return nil
}
