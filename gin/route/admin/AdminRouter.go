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
		adminGroup.GET("/addUser", adminControler.AddUser)
		adminGroup.GET("/deleteUser", adminControler.DeleteUser)
		adminGroup.GET("/updateUser", adminControler.UpdateUser)
		adminGroup.GET("/selectUser", adminControler.SelectUser)
		adminGroup.GET("/firstUser", adminControler.FirstUser)
		adminGroup.GET("/SetRedis", adminControler.SetRedis)
	}
}
