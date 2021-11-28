package config

import (
	"github.com/spf13/viper"
)

type Server struct {
	JWT    *JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis  *Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	System *System `mapstructure:"system" json:"system" yaml:"system"`
	Mysql  *Mysql  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	//	Captcha *Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	//	// gorm
	//	Pgsql Pgsql `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	//	// oss
	//	Local      Local      `mapstructure:"local" json:"local" yaml:"local"`
	//	Qiniu      Qiniu      `mapstructure:"qiniu" json:"qiniu" yaml:"qiniu"`
	//	AliyunOSS  AliyunOSS  `mapstructure:"aliyun-oss" json:"aliyunOSS" yaml:"aliyun-oss"`
	//	TencentCOS TencentCOS `mapstructure:"tencent-cos" json:"tencentCOS" yaml:"tencent-cos"`
	//	Zap        Zap        `mapstructure:"zap" json:"zap" yaml:"zap"`
}

var Conf *Server

func NewConfig() {
	vp := viper.New()
	//设置读取的配置文件
	vp.SetConfigName("config.yaml")
	//添加读取的配置文件路径
	vp.AddConfigPath("./")
	//设置配置文件类型
	vp.SetConfigType("yaml")
	// 配置文件名称(无扩展名)
	viper.SetConfigName("config")

	// 读取yaml 是否存在
	if err := vp.ReadInConfig(); err != nil {
		panic(err)
	}

	Conf = &Server{
		JWT:    jwt(vp),
		Redis:  redis(vp),
		Mysql:  mysql(vp),
		System: system(vp),
	}

}

func jwt(vp *viper.Viper) *JWT {
	jwt := &JWT{
		SigningKey:  vp.GetString("jwt.signing-key"),
		ExpiresTime: vp.GetInt64("jwt.expires-time"),
		BufferTime:  vp.GetInt64("jwt.buffer-time"),
		Issuer:      vp.GetString("jwt.issuer"),
	}
	return jwt
}

func redis(vp *viper.Viper) *Redis {
	redis := &Redis{
		DB:       vp.GetInt("redis.db"),
		Addr:     vp.GetString("redis.addr"),
		Password: vp.GetString("redis.password"),
	}
	return redis
}

func system(vp *viper.Viper) *System {
	redis := &System{
		Env:          vp.GetString("system.env"),
		Addr:         vp.GetString("system.addr"),
		DbType:       vp.GetString("system.dbType"),
		OssType:      vp.GetString("system.ossType"),
		LimitCountIP: vp.GetInt("system.ipLimitCount"),
		LimitTimeIP:  vp.GetInt("system.ipLimitTime"),
	}
	return redis
}

func mysql(vp *viper.Viper) *Mysql {
	mysql := &Mysql{
		Path:     vp.GetString("mysql.path"),
		Port:     vp.GetString("mysql.port"),
		DbName:   vp.GetString("mysql.DbName"),
		UserName: vp.GetString("mysql.UserName"),
		Password: vp.GetString("mysql.Password"),
	}
	return mysql
}
