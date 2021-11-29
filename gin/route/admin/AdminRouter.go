package routeAdmin

import (
	"github.com/gin-gonic/gin"
	adminControler "hb_gin/controller/admin"
)

type AdminRouter struct {
}

func (s *AdminRouter) InitAdmin(Router *gin.RouterGroup) {
	adminGroup := Router.Group("admin")
	var adminController = adminControler.AdminController{}
	{
		adminGroup.GET("/ping", adminController.GetPing)
		adminGroup.GET("/addUser", adminController.AddUser)
		adminGroup.GET("/deleteUser", adminController.DeleteUser)
		adminGroup.GET("/updateUser", adminController.UpdateUser)
		adminGroup.GET("/selectUser", adminController.SelectUser)
		adminGroup.GET("/firstUser", adminController.FirstUser)
		adminGroup.GET("/SetRedis", adminController.SetRedis)
	}
}
