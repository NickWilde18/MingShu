package auth

import (
	"embed"
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
)

//go:embed dist/login.html
var tplContent string

//go:embed dist
var folder embed.FS

// LoginHandler 统一处理登录页面和静态资源
func Login(r *ghttp.Request) {
	g.Log().Infof(r.Context(), "LoginHandler: %s", r.URL.Path)

	// 如果是精确访问 /auth/login，返回登录页面
	if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/login/" {
		// 来一个随机state
		state := grand.S(32)
		r.Session.Set("state", state)

		if len(tplContent) == 0 {
			r.Response.WriteStatusExit(http.StatusInternalServerError, "无法从资源中读取登录页面模板")
		}
		if err := r.Response.WriteTplContent(tplContent); err != nil {
			r.Response.WriteStatusExit(http.StatusInternalServerError, gerror.Wrap(err, "写入登录页面失败"))
		}
		return
	}

	// 否则处理静态资源
	filePath := r.URL.Path[len("/auth/login/"):]
	g.Log().Infof(r.Context(), "filePath: %s", filePath)
	content, err := folder.ReadFile("dist/" + filePath)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusNotFound, "文件不存在")
	}
	r.Response.Write(content)
}
