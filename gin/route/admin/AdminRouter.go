package routeAdmin

import (
	"github.com/gin-gonic/gin"
	adminControler "hb_gin/controller/admin"
)

type AdminRouter struct {
}

func (s *AdminRouter) InitAdmin(Router *gin.RouterGroup) {
	adminGroup := Router.Group("admin")
	{
		adminGroup.GET("/ping", adminControler.GetPing)
	}
}
