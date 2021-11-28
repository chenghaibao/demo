package viper

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"os"
	"testing"
)

// https://www.liwenzhou.com/posts/Go/viper_tutorial/ 网址
type CompanyInfomation struct {
	Name                 string
	MarketCapitalization int64
	EmployeeNum          int64
	Department           []interface{}
	IsOpen               bool
}

type YamlSetting struct {
	TimeStamp         string
	Address           string
	Postcode          int64
	CompanyInfomation CompanyInfomation
}

func Test(t *testing.T) {
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("test.yaml")
	//添加读取的配置文件路径
	v.AddConfigPath("./")
	//设置配置文件类型
	v.SetConfigType("yaml")
	// 配置文件名称(无扩展名)
	viper.SetConfigName("config")
	// 读取yaml 是否存在
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}

	fmt.Printf(
		`
		TimeStamp:%s
		CompanyInfomation.Name:%s
		CompanyInfomation.Department:%s `,
		v.Get("TimeStamp"),
		v.Get("CompanyInfomation.Name"),
		v.Get("CompanyInfomation.Department"),
	)

	fmt.Println("\n","++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	// 结构题
	parseYaml(v)

	// flag 输入参数
	flag()

	// 监听
	fmt.Println("\n","++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	watch(v)

	// env
	fmt.Println("\n","++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	env()

	// 支持etcd远程配置
}

// 监听配置文件
func watch(v *viper.Viper){
	// 监听
	v.WatchConfig()

	// 监听改变文件
	v.OnConfigChange(func(e fsnotify.Event) {
		// 配置文件发生变更之后会调用的回调函数
		fmt.Println("Config file changed:", e.Name)
		fmt.Println(v.Get("TimeStamp"))
	})

	//time.Sleep(time.Second * 10)
}

func flag(){
	// 案列
	// go run viperBase_test.go --hostAddress=192.192.1.10 --port=9000
	pflag.String("hostAddress", "127.0.0.1", "Server running address")
	pflag.Int64("port", 8080, "Server running port")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Printf("hostAddress :%s , port:%s", viper.GetString("hostAddress"), viper.GetString("port"))
}

func env(){
	viper.SetEnvPrefix("spf") // 将自动转为大写
	viper.BindEnv("id")
	os.Setenv("SPF_ID", "13") // 通常是在应用程序之外完成的
	id := viper.Get("id")
	fmt.Println(id)
}

func parseYaml(v *viper.Viper) {
	// 配置转结构体
	var yamlObj YamlSetting
	if err := v.Unmarshal(&yamlObj); err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Println(yamlObj)
}
