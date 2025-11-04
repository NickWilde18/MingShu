package uniGf

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func EnsureQP(ctx context.Context, upn string) error {
	response := client.PostVar(ctx, "/quotaPool/ensure", g.Map{
		"upn": upn,
	})
	if response.IsNil() || response.IsEmpty() {
		return gerror.New("个人配额池初始化：内部请求没有响应或返回内容为空")
	}
	content := response.Map()
	if !content["success"].(bool) {
		return gerror.New("个人配额池存在性检查失败：" + content["message"].(string))
	}
	return nil
}
