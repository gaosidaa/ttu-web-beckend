package logRecorder

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
)

var systemlogger, eventlogger *log.Logger

func PrintmynameTest() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func Log_test() {
	print("\n\n\nstart log test \n")
	fmt.Println("Hello, world!")
	log.Println("log test")

	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	var s1 string
	s1 = "output sample"
	log.Println(s1)
	//time.Sleep(3 * time.Second)
	log.Println("log test")
	CheckFile(s1)
	CreateFile()
	println("-------------------------")

	systemlogger = InitLog("System")
	eventlogger = InitLog("Event")
	println(PrintmynameTest())
	//LogOutput(eventlogger, "output test")
	jsonTest()
}

/*
func Foo() {
	fmt.Printf("我是 %s, %s 在调用我!\n", printMyName(), printCallerName())
	Bar()
}
func Bar() {
	fmt.Printf("我是 %s, %s 又在调用我!\n", printMyName(), printCallerName())
}
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
*/

type User struct {
	ID   string
	Name string
}

func jsonTest() {
	u := User{ID: "user001", Name: "tom"}
	jsonU, _ := json.Marshal(u)
	fmt.Printf("jsonU: %T\n", jsonU)
	fmt.Println(string(jsonU))
	LogOutputJson(eventlogger, jsonU)
}
