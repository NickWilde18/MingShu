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
	content := response.Map()
	if !content["success"].(bool) {
		return gerror.New("个人配额池存在性检查失败：" + content["message"].(string))
	}
	return nil
}