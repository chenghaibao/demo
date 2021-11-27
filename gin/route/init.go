package route

import (
	"github.com/gin-gonic/gin"
	"hb_gin/middleware"
	routeAdmin "hb_gin/route/admin"
	routeIndex "hb_gin/route/index"
)

type RouterGroup struct {
	Admin routeAdmin.AdminRouter
	Index routeIndex.IndexRouter
}

func Routers() *gin.Engine {
	gin := gin.Default()
	RouterGroupApp := new(RouterGroup)
	gin.Use(middleware.Auth())
	{
		RouterGroupApp.Admin.InitAdmin(gin.Group(""))
	}
	gin.Group("")
	{
		RouterGroupApp.Index.InitIndex(gin)
	}
	return gin
}
