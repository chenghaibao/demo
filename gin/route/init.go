package route

import (
	"github.com/gin-gonic/gin"
	"hb_gin/middleware"
	routeAdmin "hb_gin/route/admin"
	routeIndex "hb_gin/route/index"
)

var Gin *gin.Engine

type RouterGroup struct {
	Admin routeAdmin.AdminRouter
	Index routeIndex.IndexRouter
}

func Routers() *gin.Engine {
	Gin = gin.Default()
	RouterGroupApp := new(RouterGroup)
	Gin.Use(middleware.Auth())
	{
		RouterGroupApp.Admin.InitAdmin(Gin.Group(""))
	}
	Gin.Group("")
	{
		RouterGroupApp.Index.InitIndex(Gin)
	}
	return Gin
}
