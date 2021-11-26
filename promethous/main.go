package main

// https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#example-CounterVec
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/shirou/gopsutil/mem"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	DiskPercent = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "memeory_percent",
		Help: "memeory use percent",
	},
		[]string{"percent"},
	)

	Cpu = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "cpu_percent",
		Help: "use cpu",
	},
		[]string{"percent"},
	)

	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_code",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"code", "method"},
	)

	// SalarySummary 百分比  接口百分比  接口成功率  响应时间等等
	SalarySummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "beijing_salary",
		Help:       "the relationship between salary and population of beijing city",
		Objectives: map[float64]float64{0.5: 0.05, 0.8: 0.01, 0.9: 0.01, 0.95: 0.001},
	})
	// beijing_salary  beijing_salary_count  beijing_salary_sum

	TemperatureHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "beijing_temperature",
		Help:    "The temperature of the beijing",
		Buckets: prometheus.LinearBuckets(0, 10, 3),
	})
	// beijing_histogram_bucket  beijing_histogram_count  beijing_histogram_sum
)

func main() {
	//初始化日志服务
	logger := log.New(os.Stdout, "[Memory]", log.Lshortfile|log.Ldate|log.Ltime)
	// 启动web服务，监听8080端口
	go func() {
		logger.Println("ListenAndServe at:0.0.0.0:8973")
		err := http.ListenAndServe("127.0.0.1:8973", nil)

		if err != nil {
			logger.Fatal("ListenAndServe: ", err)
		}
	}()

	//初始一个http handler
	http.Handle("/metrics", promhttp.Handler())
	prometheus.MustRegister(DiskPercent)
	prometheus.MustRegister(Cpu)
	prometheus.MustRegister(httpReqs)
	prometheus.MustRegister(SalarySummary)
	prometheus.MustRegister(TemperatureHistogram)

	//httpReqs.DeleteLabelValues("200", "GET")
	//// Same thing with the more verbose Labels syntax.
	//httpReqs.Delete(prometheus.Labels{"method": "GET", "code": "200"})

	for {
		v, err := mem.VirtualMemory()
		logger.Println("memory----", v)
		if err != nil {
			logger.Println("get memeory use percent error:%s", err)
		}
		requestSecond := rand.Intn(100)

		// 百分比
		SalarySummary.Observe(float64(requestSecond))
		usedPercent := v.UsedPercent

		// gauge
		DiskPercent.WithLabelValues("usedMemory").Set(usedPercent)
		Cpu.WithLabelValues("usedCpu").Set(usedPercent)

		// counter
		httpReqs.WithLabelValues("404", "POST").Add(1)
		httpReqs.WithLabelValues("200", "GET").Inc()

		time.Sleep(time.Second * 2)

		TemperatureHistogram.Observe(v.UsedPercent)
		// 百度案列
		//var temperature = [10]float64{1, 4, 5, 10, 14, 15, 20, 25, 11, 30}
		//for i := 0; i < len(temperature); i++ {
		//fmt.Printf("insert number: %f \n", temperature[i])
		//}
	}

}
