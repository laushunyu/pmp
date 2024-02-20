package main

import (
	_ "github.com/laushunyu/pmp/internal/packed"

	_ "github.com/laushunyu/pmp/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"github.com/laushunyu/pmp/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
