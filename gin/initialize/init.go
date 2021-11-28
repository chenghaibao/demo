package initialize

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"hb_gin/common"
	"hb_gin/config"
	"hb_gin/initialize/mysql"
	"hb_gin/initialize/redis"
	"hb_gin/route"
	"time"
)

type server interface {
	ListenAndServe() error
}

func Run() {
	// 加载配置
	config.NewConfig()
	// 加载日志
	common.NewLogger()
	// 加载定时器
	//common.NewCron()
	// 加载mysql
	mysql.NewMysql()
	// 加载redis
	redis.NewRedis()
	// 初始化路由
	gin := route.Routers()
	// 开始玩耍了
	runServer(gin)
}

func runServer(gin *gin.Engine) {
	s := initServer(":"+config.Conf.System.Addr,gin)
	// 写入错误日志
	common.Log.WriteError(s.ListenAndServe().Error())
}


func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
