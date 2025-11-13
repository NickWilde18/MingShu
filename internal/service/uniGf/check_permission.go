package uniGf

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// CheckPermission 检查用户权限
func CheckPermission(ctx context.Context, sub string, obj string, act string) error {
	response := client.PostVar(ctx, "/auth/check", g.Map{
		"sub": sub,
		"obj": obj,
		"act": act,
	})
	if response.IsNil() || response.IsEmpty() {
		return gerror.New("权限检查：内部请求没有响应或返回内容为空")
	}
	content := response.Map()
	if !content["success"].(bool) {
		return gerror.New("权限检查失败：" + content["message"].(string))
	}
	if !content["data"].(g.Map)["allow"].(bool) {
		return gerror.Newf("你没有权限使用 [%s] %s。请联系管理员。", obj, act)
	}
	return nil
}
