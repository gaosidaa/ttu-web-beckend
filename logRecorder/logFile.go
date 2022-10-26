package logRecorder

import (
	"log"
	"os"
	"time"
)

func CheckFile(Filename string) bool {
	var exist = true
	timeString := time.Now().Format("2006-01-02")
	//file := "./logData/" + timeString + ".log"
	file := "./logData/" + Filename + "-" + timeString + ".log"
	Filename = file
	if _, err := os.Stat(Filename); os.IsNotExist(err) {
		exist = false
		if err != nil {
			//fmt.Println("Check file error!")
		}
	}
	return exist
}

/*
func CreateFile() bool {
	timeString := time.Now().Format("2006-01-02")
	file := "./logRecorder/LogData/" + timeString + ".log"
	println(file)
	return true
}

*/

//写入文件
/*
func Logfile(Log string) {
	var f1 *os.File
	var err1 error

	Filenames := "./launch.log" //也可将name作为参数传进来

	if CheckFile(Filenames) { //如果文件存在
		f1, err1 = os.OpenFile(Filenames, os.O_APPEND|os.O_WRONLY, 0666) //打开文件,第二个参数是写入方式和权限
		if err1 != nil {
			fmt.Println("文件存在，已打开")
		}
	} else {
		f1, err1 = os.Create(Filenames) //创建文件
		if err1 != nil {
			fmt.Println("创建文件失败")
		}
	}
	_, err1 = io.WriteString(f1, Log+"\n") //写入文件(字符串)
	if err1 != nil {
		fmt.Println(err1)
	}
	//fmt.Printf("写入 %d 个字节\n", n)

	return
}

*/

/*--------------------------------------*/
/*-----------new version----------------*/

func InitLog(name string) *log.Logger {
	//创建文件对象, 日志的格式为当前时间2006-01-02 15:04:05.log;据说是golang的诞生时间，固定写法
	timeString := time.Now().Format("2006-01-02")
	//file := "./" + timeString + ".log"
	file := "./logData/" + name + "-" + timeString + ".log"

	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}

	//创建一个Logger, 参数1：日志写入的文件, 参数2：每条日志的前缀；参数3：日志属性
	logpointer := log.New(logFile, "logpre_", log.Lshortfile)

	//Prefix返回前缀，Flags返回Logger的输出选项属性值
	//fmt.Printf("\n\n创建时前缀为:%s\n创建时输出项属性值为:%d\n", logpointer.Prefix(), logpointer.Flags())

	//SetFlags 重新设置输出选项
	logpointer.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	logpointer.SetFlags(log.Ldate | log.Ltime)

	//重新设置输出前缀
	logpointer.SetPrefix(name)

	//获取修改后的前缀和输出项属性值
	//fmt.Printf("修改后前缀为:%s\n修改后输出项属性值为:%d\n", logpointer.Prefix(), logpointer.Flags())

	logpointer.Output(2, "打开日志")

	return logpointer
	//return log.New(logFile, "logpre_", log.Lshortfile)
}

func logOutput(logpointer *log.Logger, content string) {
	logpointer.Output(2, content)
}

/*

func logOutputJson(logpointer *log.Logger, content []uint8) {
	logpointer.Output(2, string(content))
}

func outputlog(logpointer *log.Logger, content string, kind int) *log.Logger {
	//kind : 1 System
	//kind : 2 Event
	var fileKind string
	switch {
	case kind == 1:
		fileKind = "System"
	case kind == 2:
		fileKind = "Event"
	default:
		fileKind = "Event"
	}
	//println(CheckFile(fileKind))
	if CheckFile(fileKind) == false {
		logpointer = InitLog(fileKind)
	}
	logOutput(logpointer, content)
	return logpointer
}

*/
