package routeIndex

import (
	"github.com/gin-gonic/gin"
	adminControler "hb_gin/controller/admin"
)

type IndexRouter struct {
}

func (s *IndexRouter) InitIndex(gin *gin.Engine) {
	var adminController = adminControler.AdminController{}
	gin.GET("/ping", adminController.GetPing)
}
