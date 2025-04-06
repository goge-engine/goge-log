package gogeLog

import (
	"errors"
	"testing"

	gogeBasis "github.com/goge-engine/goge-basis"
)

func TestGogeLog(t *testing.T) {
	var logger GogeLogClass
	var errorTest error
	errorTest = errorForTest()

	logger.InitGogeLog("./gogeLog.log")
	logger.Log("[gogeLog] test log")
	logger.Error(errorTest)
	logger.Warning(gogeBasis.ReturnWarningForTest())
	logger.DevelopmentLog("gogeLogTest", "test development log")

	defer logger.Close()
}

func errorForTest() error {
	return errors.New("test error")
}
