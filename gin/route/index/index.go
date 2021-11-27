package routeIndex

import (
	"github.com/gin-gonic/gin"
	adminControler "hb_gin/controller/admin"
)

type IndexRouter struct {
}

func (s *IndexRouter) InitIndex(gin *gin.Engine) {
	gin.GET("/ping", adminControler.GetPing)
}
