package config

import (
	"fmt"
	"github.com/spf13/viper"
	"hb_distributeStorage/utils"
)

type Conf struct {
	Host    string `json:"host,omitempty"`
	Port    string `json:"port"`
	Cluster string `json:"cluster"`
}

var Config *Conf

func NewConfig() *Conf {
	v := viper.New()
	//设置读取的配置文件
	v.SetConfigName("config.yaml")
	//添加读取的配置文件路径
	v.AddConfigPath("./config/")
	//设置配置文件类型
	v.SetConfigType("yaml")
	// 配置文件名称(无扩展名)
	viper.SetConfigName("config")
	// 读取yaml 是否存在
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("err:%s\n", err)
	}
	var Cluster string
	if v.Get("address") == nil {
		Cluster = utils.Strval(v.Get("host")) + ":" + utils.Strval(v.Get("port"))
	} else {
		Cluster = utils.Strval(v.Get("address"))
	}
	Config = &Conf{
		Host:    utils.Strval(v.Get("host")),
		Cluster: Cluster,
		Port:    utils.Strval(v.Get("port")),
	}
	return Config
}
