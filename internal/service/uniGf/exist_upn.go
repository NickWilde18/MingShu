package uniGf

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	m "uniauth-gateway/internal/middlewares"
	"uniauth-gateway/internal/consts"
)

// ExistUPN 检查 UPN 是否存在
func ExistUPN(ctx context.Context, upn string) error {
	response := client.GetVar(ctx, "/userinfos", g.Map{
		"upn": upn,
	})
	if response.IsNil() || response.IsEmpty() {
		return gerror.New("校验用户信息：内部请求没有响应或返回内容为空")
	}
	content := response.Map()
	if !content["success"].(bool) {
		r := g.RequestFromCtx(ctx)
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeForbidden,
			ShowDetail: true,
			Detail: `无法查找到你的AD域信息。这通常不是你的问题。
如果您是新入职/入学的人员，请在1天后再登录。
如您已经入职/入学1天以上，或想立即进入平台，请联系管理员手动同步AD域数据库。

AD domain information could not be found. This is usually not your issue.
If you are a new employee/student, please try logging in again after 1 day.
If you have been employed/enrolled for more than 1 day, or wish to access the platform immediately, please contact the administrator to manually synchronize the AD domain database.`,
		})
		r.Exit()
	}
	return nil
}
