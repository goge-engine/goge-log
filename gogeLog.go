package gogeLog

import (
	"fmt"
	"log"
	"os"

	gogeBasis "github.com/goge-engine/goge-basis"
)

type GogeLogClass struct {
	Logger *log.Logger
	File   *os.File
}

func (this *GogeLogClass) InitGogeLog(Path string) {
	file, err := os.OpenFile(Path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("[gogeLog] ERROR IN LOG INIT:", err)
	}
	this.File = file
	this.Logger = log.New(this.File, "[gogeLog] ", log.Ldate|log.Ltime|log.Lshortfile)
	this.Logger.SetOutput(this.File)
	this.Logger.Println("[gogeLog] INIT SUCCESS")
}

func (this *GogeLogClass) Log(logContext string) {
	this.Logger.Println(logContext)
}

func (this *GogeLogClass) Error(errorContext error) {
	if errorContext != nil {
		this.Logger.Println("[error]", errorContext)
	}
}

func (this *GogeLogClass) Warning(warningContext gogeBasis.Warning) {
	if warningContext != nil {
		this.Logger.Println("[warning]", warningContext.Warning())
	}
}

func (this *GogeLogClass) DevelopmentLog(logHead string, logContext string) {
	this.Logger.Println("["+logHead+"]", logContext)
}

func (this *GogeLogClass) ErrorAndWarningProcess(errorContext error, warningContext gogeBasis.Warning) {
	if errorContext != nil {
		this.Error(errorContext)
	}
	if warningContext != nil {
		this.Warning(warningContext)
	}
	return
}

func (this *GogeLogClass) Close() {
	this.File.Close()
}
