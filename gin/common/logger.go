package common

import (
	"github.com/op/go-logging"
	"io"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Log *Logging

type Logging struct {
	logging.Logger
}

// 继承新的
func (l *Logging) WriteError(args string) {
	// 存入文件
	writeLog("err.log", args)
	l.Error(args)
}

func NewLogger() *Logging {
	logger := logging.MustGetLogger("hb")

	format := logging.MustStringFormatter(
		`%{time:2006-01-02 15:04:05.000} %{longfunc}[%{shortfile}] ▶ [%{level:.4s}] %{message}`,
	)
	// os.Stderr  获取默认的日志
	backend := logging.NewLogBackend(os.Stderr, "hb", 0)
	//// 输出日志的格式
	backendFormatter := logging.NewBackendFormatter(backend, format)
	//// Only errors and more severe messages should be sent to backend1
	backendLeveled := logging.AddModuleLevel(backend)
	backendLeveled.SetLevel(logging.ERROR, "")

	// Set the backends to be used.
	logging.SetBackend(backendLeveled, backendFormatter)

	Log = &Logging{
		*logger,
	}

	return Log
}

// 严重错误写入文件
func writeLog(fileName, msg string) error {
	path := "log/" + time.Now().Format("20060102") + "/"
	if !isExist(path) {
		return createDir(path)
	}
	var (
		err error
		f   *os.File
	)
	f, err = os.OpenFile(path+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	time := time.Now().Format("2006-01-02 15:04:05")
	_, file, line, _ := runtime.Caller(1)
	str := time + "|" + file + "|" + strconv.Itoa(line) + "|" + msg + "\n"
	_, err = io.WriteString(f, str)
	defer f.Close()
	return err
}

//CreateDir  文件夹创建
func createDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

//IsExist  判断文件夹/文件是否存在  存在返回 true
func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
