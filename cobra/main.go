package main

import (
	"log"
	"main.go/cmd"
)
// https://juejin.cn/post/6924541628031959047
func main()  {
	// 日志等级
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	cmd.Execute()
}
