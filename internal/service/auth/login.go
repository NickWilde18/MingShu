package auth

import (
	_ "embed"
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
)

//go:embed login.html
var tplContent string

func Login(r *ghttp.Request) {
	// 来一个随机state
	state := grand.S(32)
	r.Session.Set("state", state)

	if len(tplContent) == 0 {
		r.Response.WriteStatusExit(http.StatusInternalServerError, "无法从资源中读取登录页面模板")
	}
	if err := r.Response.WriteTplContent(tplContent); err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		r.Response.Write(gerror.Wrap(err, "写入登录页面失败").Error())
	}
}
