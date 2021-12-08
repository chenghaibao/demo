package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt验证
		//token := c.GetHeader("Authorization")
		//userToken, err := utils.GetToken(token)
		//isSuccess := checkJwt(userToken, err)
		//if isSuccess {
		//	c.Next()
		//} else {
		//	// 错误提示
		//	c.Abort()
		//}

		// 测试代码  请自己删除
		success := "success"
		if success == "success" {
			c.Next()
		}
		c.Abort()
	}
}

func checkJwt(token *jwt.Token, err error) bool {
	if token.Valid {
		return false
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
