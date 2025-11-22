package auth

import (
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/golang-jwt/jwt/v5"

	"uniauth-gateway/internal/consts"
	m "uniauth-gateway/internal/middlewares"
	"uniauth-gateway/internal/service/uniGf"
)

func Callback(r *ghttp.Request) {
	state := r.Get("state").String()
	oState, err := r.Session.Get("state")
	if removeErr := r.Session.Remove("state"); removeErr != nil {
		g.Log().Error(r.Context(), gerror.Wrap(removeErr, "清除 Session 失败"))
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeCallbackFailed,
			Detail:    removeErr.Error(),
		})
	}
	if err != nil || oState.IsNil() || oState.String() != state {
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeInvalidToken,
			Detail:    "State 不存在或校验不通过",
		})
	}
	code := r.Get("code").String()
	ctx := r.Context()

	response := g.Client().ContentType("application/x-www-form-urlencoded; charset=utf-8").PostVar(
		ctx,
		g.Cfg().MustGet(ctx, "sso.tokenUrl").String(),
		g.Map{
			"client_id":     g.Cfg().MustGet(ctx, "sso.clientId").String(),
			"code":          code,
			"redirect_uri":  g.Cfg().MustGet(ctx, "sso.redirectUri").String(),
			"grant_type":    "authorization_code",
			"client_secret": g.Cfg().MustGet(ctx, "sso.clientSecret").String(),
		},
	).Map()

	_, ok := response["error"]
	if ok {
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeCallbackFailed,
			Detail: fmt.Sprintf(
				`因为以下原因，SSO认证失败，请重试。
SSO authentication failed due to the following reason. Please try again.
错误/Error: %s
错误描述/Error Description: %s`,
				response["error"],
				response["error_description"],
			),
		})
	}

	accessToken, ok := response["access_token"]
	if !ok {
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeCallbackFailed,
			Detail: `SSO认证信息返回信息中没有 Access Token。可能是登录过于频繁，请稍后再试。
The SSO authentication response does not contain an Access Token. This may be due to excessive login attempts; please try again later.`,
		})
	}

	parser := jwt.NewParser()
	token, _, err := parser.ParseUnverified(accessToken.(string), &jwt.MapClaims{})
	if err != nil {
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeCallbackFailed,
			Detail: `JWT 解析失败。
The JWT parsing failed.`,
		})
	}
	upn, ok := (*(token.Claims.(*jwt.MapClaims)))["upn"]
	if !ok {
		m.RenderError(r, m.ErrorInfo{
			ErrorCode: consts.ErrCodeCallbackFailed,
			Detail: `JWT 中没有 upn 字段。
The JWT does not contain the upn field.`,
		})
	}

	// 先看upn存不存在
	uniGf.ExistUPN(ctx, upn.(string))
	// 然后必须 Ensure QP
	uniGf.EnsureQP(ctx, upn.(string))
	// 再看upn有没有权限进入
	uniGf.CheckPermission(ctx, upn.(string), "platform", "access")

	// 记录 Session
	r.Session.RegenerateId(true)
	r.Session.Set("user_id", upn.(string))
	r.Response.RedirectTo("/chat/")
}
