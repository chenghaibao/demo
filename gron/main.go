package main

import (
	"context"
	"fmt"
	"github.com/roylee0704/gron"
	"os"
	"time"
)

var cron *gron.Cron
var job *PrintJob

func NewCron() *gron.Cron {
	cron = gron.New()
	job = &PrintJob{"16", "hb"}
	return cron
}

type PrintJob struct {
	Age  string
	Name string
}

func (r *PrintJob) Run() {
	fmt.Println(r.Age, r.Name)
}

func main() {
	NewCron()

	// 设定结束参数
	endChan := make(chan bool, 1)
	ctx, cancel := context.WithCancel(context.Background())

	// 结构体
	job.Age = "17"
	job.Name = "HB"
	cron.Add(gron.Every(1*time.Second), job)
	cron.AddFunc(gron.Every(1*time.Second), ping)

	//结束运行通道
	cron.AddFunc(gron.Every(3*time.Second), func() {
		endChan <- true
	})

	// 结束运行ctx 上下文
	cron.AddFunc(gron.Every(2*time.Second), func() {
		cancel()
	})

	// 定时器开始
	cron.Start()

	//test(endChan)
	str := "sadsadsa"
	testString(&str)
	for {
		select {
		case test := <-endChan:
			fmt.Println("test111", test)
			cron.Stop()
			os.Exit(1)
		case <-ctx.Done():
			fmt.Println("结束了")
			cron.Stop()
			os.Exit(1)
		}
	}

}

func ping() {
	fmt.Println("21321")
}

func test(test chan<-bool){
	test<-true
}

func testString(str *string){
	fmt.Println(str)
}
