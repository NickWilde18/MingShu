package uniGf

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"uniauth-gateway/internal/consts"
	m "uniauth-gateway/internal/middlewares"
)

func EnsureQP(ctx context.Context, upn string) {
	response := client.PostVar(ctx, "/quotaPool/ensure", g.Map{
		"upn": upn,
	})
	if response.IsNil() || response.IsEmpty() {
		m.RenderError(g.RequestFromCtx(ctx), m.ErrorInfo{
			ErrorCode: consts.ErrCodeBadGateway,
			Detail:    `个人配额池初始化：内部请求没有响应或返回内容为空。
The personal quota pool initialization: the internal request does not respond or the returned content is empty.`,
		})
	}
	content := response.Map()
	if !content["success"].(bool) {
		m.RenderError(g.RequestFromCtx(ctx), m.ErrorInfo{
			ErrorCode: consts.ErrCodeBadGateway,
			Detail:    "个人配额池存在性检查失败：" + content["message"].(string) + `
The personal quota pool existence check failed: ` + content["message"].(string),
		})
	}
}
