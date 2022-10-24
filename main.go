package main

import (
	_ "ttu-backend/internal/packed"
	"ttu-backend/logRecorder"
)

func main() {
	//cmd.Main.Run(gctx.New())
	println(logRecorder.PrintmynameTest())
	logRecorder.Log_test()
}
