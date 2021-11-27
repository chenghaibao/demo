package main

import (
	"github.com/op/go-logging"
	"hb_go_logging/utils"
)

type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {
	utils.NewLogger()
	utils.Log.Debugf("debug %s", Password("secret"))
	utils.Log.Notice("notice")
	utils.Log.Warning("warning")
	utils.Log.WriteError("err")
	utils.Log.Critical("crit")
	utils.Log.Info("12321")
}

