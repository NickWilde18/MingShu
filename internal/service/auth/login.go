package auth

import (
	"embed"
	"net/http"
	"path/filepath"
	"strings"
	"uniauth-gateway/internal/consts"
	m "uniauth-gateway/internal/middlewares"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
)

//go:embed login/login.html
var tplContentLogin string

//go:embed login
var folderLogin embed.FS

//go:embed  login-legacy/login-legacy.html
var tplContentLegacy string

//go:embed login-legacy
var folderLagacy embed.FS

func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".css":
		return "text/css"
	case ".js":
		return "application/javascript"
	case ".png":
		return "image/png"
	case ".html":
		return "text/html"
	default:
		return "application/octet-stream"
	}
}

// LoginHandler 统一处理登录页面和静态资源
func Login(r *ghttp.Request) {
	g.Log().Infof(r.Context(), "LoginHandler: %s", r.URL.Path)

	// 如果是精确访问 /auth/login，返回登录页面
	if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/login/" {
		// 来一个随机state
		state := grand.S(32)
		r.Session.Set("state", state)

		if len(tplContentLogin) == 0 {
			m.RenderError(r, m.ErrorInfo{
				ErrorCode: consts.ErrCodeInternalServer,
				Detail: "无法从资源中读取登录页面模板。<br/>The login page template cannot be read from the resource.",
			})
		}
		if err := r.Response.WriteTplContent(tplContentLogin); err != nil {
			m.RenderError(r, m.ErrorInfo{
				ErrorCode: consts.ErrCodeInternalServer,
				Detail: "写入登录页面失败。<br/>The login page template cannot be written to the resource.",
			})
		}
		return
	}

	// 否则处理静态资源
	filePath := r.URL.Path[len("/auth/login/"):]
	g.Log().Infof(r.Context(), "filePath: %s", filePath)
	content, err := folderLogin.ReadFile("login/" + filePath)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusNotFound, "文件不存在")
	}

	// 设置正确的 Content-Type
	contentType := getContentType(filePath)
	r.Response.Header().Set("Content-Type", contentType)

	r.Response.Write(content)
}

// LoginLegacyHandler 统一处理登录页面和静态资源（兼容页面版本）
// 逻辑和上面一模一样，但为了方便阅读和维护，分成了两个函数
func LoginLegacy(r *ghttp.Request) {
	g.Log().Infof(r.Context(), "LoginLegacyHandler: %s", r.URL.Path)

	// 如果是精确访问 /auth/login，返回登录页面
	if r.URL.Path == "/auth/login-legacy" || r.URL.Path == "/auth/login-legacy/" {
		// 来一个随机state
		state := grand.S(32)
		r.Session.Set("state", state)

		if len(tplContentLegacy) == 0 {
			m.RenderError(r, m.ErrorInfo{
				ErrorCode: consts.ErrCodeInternalServer,
				Detail: "无法从资源中读取登录页面模板。<br/>The login page template cannot be read from the resource.",
			})
		}
		if err := r.Response.WriteTplContent(tplContentLegacy); err != nil {
			m.RenderError(r, m.ErrorInfo{
				ErrorCode: consts.ErrCodeInternalServer,
				Detail: "写入登录页面失败。<br/>The login page template cannot be written to the resource.",
			})
		}
		return
	}

	// 否则处理静态资源
	filePath := r.URL.Path[len("/auth/login-legacy/"):]
	g.Log().Infof(r.Context(), "filePath: %s", filePath)
	content, err := folderLagacy.ReadFile("login-legacy/" + filePath)
	if err != nil {
		r.Response.WriteStatusExit(http.StatusNotFound, "文件不存在")
	}

	// 设置正确的 Content-Type
	contentType := getContentType(filePath)
	r.Response.Header().Set("Content-Type", contentType)

	r.Response.Write(content)
}
