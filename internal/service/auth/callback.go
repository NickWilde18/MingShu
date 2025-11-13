package auth

import (
	"net/http"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"

	"uniauth-gateway/internal/service/uniGf"
)

func Callback(r *ghttp.Request) {
	state := r.Get("state").String()
	oState, err := r.Session.Get("state")
	if removeErr := r.Session.Remove("state"); removeErr != nil {
		g.Log().Errorf(r.Context(), "清除 Session 失败: %v", removeErr.Error())
	}
	if err != nil || oState.IsNil() || oState.String() != state {
		r.Response.WriteStatusExit(http.StatusUnauthorized, "State 不存在或校验不通过")
	}
	code := r.Get("code").String()
	ctx := r.Context()

	response := g.Client().ContentType("application/x-www-form-urlencoded; charset=utf-8").PostVar(
		ctx,
		g.Cfg().MustGet(ctx, "sso.token_url").String(),
		g.Map{
			"client_id":     g.Cfg().MustGet(ctx, "sso.client_id").String(),
			"code":          code,
			"redirect_uri":  g.Cfg().MustGet(ctx, "sso.redirect_uri").String(),
			"grant_type":    "authorization_code",
			"client_secret": g.Cfg().MustGet(ctx, "sso.client_secret").String(),
		},
	).Map()

	_, ok := response["error"]
	if ok {
		r.Response.WriteStatusExit(http.StatusInternalServerError, gerror.Newf(
			"因为以下原因，SSO认证失败，请重试。<br/>SSO authentication failed due to the following reason. Please try again.<br/><br/>错误/Error: %s<br/><br/>错误描述/Error Description: %s",
			response["error"],
			response["error_description"],
		))
	}

	accessToken, ok := response["access_token"]
	if !ok {
		r.Response.WriteStatusExit(http.StatusInternalServerError, gerror.New(
			"SSO认证信息返回信息中没有Access Token。这通常不是你的问题，请联系管理员检查日志。<br/>Can not get your access token. This is usually not your issue. Please contact the administrator to check the logs.",
		))
	}

	parser := jwt.NewParser()
	token, _, err := parser.ParseUnverified(accessToken.(string), &jwt.MapClaims{})
	if err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, gerror.Wrap(err, "JWT 解析失败"))
	}
	upn, ok := (*(token.Claims.(*jwt.MapClaims)))["upn"]
	if !ok {
		r.Response.WriteStatusExit(http.StatusInternalServerError, gerror.New("JWT 中没有 upn 字段"))
	}

	g.Log().Infof(ctx, "upn: %s", upn)
	// 先看upn存不存在
	if err = uniGf.ExistUPN(ctx, upn.(string)); err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, err)
	}
	// 再看upn有没有权限进入
	if err = uniGf.CheckEntryPermission(ctx, upn.(string)); err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, err)
	}
	// 最后Ensure QP
	if err = uniGf.EnsureQP(ctx, upn.(string)); err != nil {
		r.Response.WriteStatusExit(http.StatusInternalServerError, err)
	}

	// 记录 Session
	r.Session.RegenerateId(true)
	r.Session.Set("user_id", upn.(string))
	r.Response.RedirectTo("/uniauth/")
}
