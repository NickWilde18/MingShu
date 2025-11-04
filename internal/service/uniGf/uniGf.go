package uniGf

import (
	"time"
	
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/net/gclient"
)

var client *gclient.Client

func init() {
	ctx := gctx.GetInitCtx()
	client = g.Client().ContentJson()
	client.SetPrefix(g.Cfg().MustGet(ctx, "uniGf.baseURL").String())
	client.SetTimeout(5 * time.Second)
}