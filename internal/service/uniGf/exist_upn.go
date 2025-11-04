package uniGf

import (
	"context"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// ExistUPN 检查 UPN 是否存在
func ExistUPN(ctx context.Context, upn string) error {
	response, err := client.Get(ctx, "/userinfos", g.Map{
		"upn": upn,
	})
	if err != nil {
		return gerror.Wrap(err, "调用用户信息API失败")
	}
	defer response.Close()

	if response.StatusCode != 200 {
		return gerror.Newf("用户信息API返回错误状态码: %d", response.StatusCode)
	}

	responseBody := response.ReadAll()
	jsonObj, err := gjson.DecodeToJson(responseBody)
	if err != nil {
		return gerror.Wrap(err, "解析API响应失败")
	}
	content := jsonObj.Map()
	g.Dump(content)
	if !content["success"].(bool) {
		return gerror.New(
			`无法查找到你的AD域信息。这通常不是你的问题。
如果您是新入职/入学的人员，请在1天后再登录。
如您已经入职/入学1天以上，或想立即进入平台，请联系管理员手动同步AD域数据库。

AD domain information could not be found. This is usually not your issue.
If you are a new employee/student, please try logging in again after 1 day.
If you have been employed/enrolled for more than 1 day, or wish to access the platform immediately, please contact the administrator to manually synchronize the AD domain database.`)
	}
	return nil
}
