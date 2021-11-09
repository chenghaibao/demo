package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
)

var (
	DiskPercent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memeory_percent",
		Help: "memeory use percent",
	},
		[]string {"percent"},
	)

	Cpu = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cpu",
		Help: "use cpu",
	},
		[]string {"cpu"},
	)
)
func main (){
	//初始化日志服务
	logger := log.New(os.Stdout, "[Memory]", log.Lshortfile | log.Ldate | log.Ltime)

	//初始一个http handler
	http.Handle("/metrics", promhttp.Handler())

	//初始化一个容器
	prometheus.MustRegister(DiskPercent)
	prometheus.MustRegister(Cpu)

	// 启动web服务，监听8080端口
	go func() {
		logger.Println("ListenAndServe at:0.0.0.0:8973")
		err := http.ListenAndServe("127.0.0.1:8973", nil)

		if err != nil {
			logger.Fatal("ListenAndServe: ", err)
		}
	}()

	//收集内存使用的百分比
	for {
		logger.Println("start collect memory used percent!")
		v, err := mem.VirtualMemory()
		if err != nil {
			logger.Println("get memeory use percent error:%s", err)
		}
		usedPercent := v.UsedPercent
		logger.Println("get memeory use percent:", usedPercent)
		DiskPercent.WithLabelValues("usedMemory").Set(usedPercent)
		time.Sleep(time.Second*2)
	}
}