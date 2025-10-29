package service

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var proxyHostMap g.Map = g.Cfg("mingshu-config").MustGet(gctx.GetInitCtx(), "proxy_host_map").Map()

func GetProxyHost(service string) (string, error) {
	proxyHost, ok := proxyHostMap[service]
	if !ok {
		return "", gerror.Newf("未找到 %s 服务。请确认请求头是否传递正确的服务名称！", service)
	}
	return proxyHost.(string), nil
}
