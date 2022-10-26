package logRecorder

import (
	"log"
	"runtime"
)

var systemlogger, eventlogger *log.Logger

func PrintmynameTest() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func PrintCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

/*
func Log_test() {
	print("\n\n\nstart log test \n")
	//fmt.Println("Hello, world!")
	//log.Println("log test")

	//log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	//log.SetFlags(log.Llongfile | log.Ltime | log.Ldate)
	//var s1 string
	//s1 = "output sample"
	//log.Println(s1)
	//time.Sleep(3 * time.Second)
	//log.Println("log test")
	//println(CheckFile("Event"))
	//CreateFile()
	println("-------------------------")

	//systemlogger = initLog("System")
	//eventlogger = initLog("Event")
	//systemlogger = outputlog(systemlogger, "test sample for new", 1)
	//eventlogger = outputlog(eventlogger, "test sample for new", 2)

	//jsonTest()

	var systemLog SystemLog
	var eventLog EventLog
	systemLog = SystemLog{InitLog("System")}
	eventLog = EventLog{InitLog("Event")}
	systemLog.Writelog("content")
	systemLog.Writelog("content")
	eventLog.Writelog("content")
	eventLog.Writelog("content")
}
*/

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

/*
type User struct {
	ID   string
	Name string
}
func jsonTest() {
	u := User{ID: "user001", Name: "tom"}
	jsonU, _ := json.Marshal(u)
	fmt.Printf("jsonU: %T\n", jsonU)
	fmt.Println(string(jsonU))
	logOutputJson(eventlogger, jsonU)

}
*/

type Log interface {
	Writelog()
}

type SystemLog struct {
	Pointer *log.Logger
}

func (sample SystemLog) Writelog(content string) {
	//println(CheckFile("System"))
	if CheckFile("System") == false {
		sample.Pointer = InitLog("System")
	}
	logOutput(sample.Pointer, PrintCallerName()+content)
}

type EventLog struct {
	Pointer *log.Logger
}

func (sample EventLog) Writelog(content string) {
	//println(CheckFile("System"))
	if CheckFile("System") == false {
		sample.Pointer = InitLog("Event")
	}
	logOutput(sample.Pointer, PrintCallerName()+content)
}
