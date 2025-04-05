package gogeLog

import (
	"fmt"
	"testing"

	gogeBasis "github.com/goge-engine/goge-basis"
)

func TestGogeLog(t *testing.T) {
	var logger GogeLogClass
	var errorTest error
	errorTest = fmt.Errorf("test error")

	logger.InitGogeLog("./gogeLog.log")
	logger.Log("[gogeLog] test log")
	logger.Error(errorTest)
	logger.Wraning(gogeBasis.Warning{Message: "test warning"})
	logger.DevelopmentLog("gogeLogTest", "test development log")

	defer logger.Close()
}
