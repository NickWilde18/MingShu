package auth

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/gogf/gf/v2/util/grand"
)

func Login(r *ghttp.Request) {
	ctx := r.Context()
	// 来一个随机state
	state := grand.S(32)
	r.Session.Set("state", state)
	
	// 从 gres 打包资源中读取模板内容
	tplContent := gres.GetContent("resource/template/login.html")
	if len(tplContent) == 0 {
		r.Response.WriteStatus(http.StatusInternalServerError)
		r.Response.Write("无法从资源中读取登录页面模板")
		return
	}
	if err := r.Response.WriteTplContent(string(tplContent),
		g.Map{
			"client_id":    g.Cfg().MustGet(ctx, "sso.client_id").String(),
			"redirect_uri": g.Cfg().MustGet(ctx, "sso.redirect_uri").String(),
			"resource":     g.Cfg().MustGet(ctx, "sso.resource").String(),
			"state":        state,
		}); err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError)
		r.Response.Write(gerror.Wrap(err, "写入登录页面失败").Error())
	}
}
