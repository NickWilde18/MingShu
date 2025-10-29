package main

import (
	_ "MingShu/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"MingShu/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
