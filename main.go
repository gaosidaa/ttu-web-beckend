package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"ttu-backend/internal/cmd"
	_ "ttu-backend/internal/packed"
)

func main() {
	cmd.Main.Run(gctx.New())
	//println(logRecorder.PrintmynameTest())
	//logRecorder.Log_test()

}
