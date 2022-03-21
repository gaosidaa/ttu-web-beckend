package main

import (
	_ "ttu-backend/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"ttu-backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
