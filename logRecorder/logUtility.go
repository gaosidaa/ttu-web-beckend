package logRecorder

import "runtime"

func PrintmynameTest() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func TestFunc() {
	println("test func")
}
