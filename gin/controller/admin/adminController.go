package adminControler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "sadsa",
	})

}
