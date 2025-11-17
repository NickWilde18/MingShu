package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"

	"uniauth-gateway/internal/cmd"
)

func main() {
	err := gtime.SetTimeZone("Asia/Shanghai")
    if err != nil {
        panic(err)
    }
	cmd.Main.Run(gctx.GetInitCtx())
}
