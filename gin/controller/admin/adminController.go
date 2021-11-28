package adminControler

import (
	"context"
	"fmt"
	mysql "hb_gin/plugin/mysql"
	redis "hb_gin/plugin/redis"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	UserModel "hb_gin/Model/User"
	AdminService "hb_gin/service/Admin"
)

func GetPing(c *gin.Context) {
	sum := AdminService.GetPing(2)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "sadsa" + strconv.Itoa(sum),
	})
}

func AddUser(c *gin.Context) {
	user := &UserModel.User{
		Nickname: "Nickname",
		Birth:    "Birth",
		Sex:      "12",
		Height:   12,
		Hobby:    "Hobby",
		Remark:   "Remark",
	}
	user.Add()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

func DeleteUser(c *gin.Context) {
	user := new(UserModel.User)
	mysql.Db.Where("id = ?", 1).Delete(user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

func UpdateUser(c *gin.Context) {
	// 更改1
	//mysql.Db.Model(&UserModel.User{}).Where("id = ?", 2).Update("nickname","test")
	// 更改2
	//var user UserModel.User
	//mysql.Db.Where("id = ?", 3).First(&user)
	//user.Nickname = "zisefeizhu"
	//user.Birth = "23"
	//mysql.Db.Debug().Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

func SelectUser(c *gin.Context) {
	var userArr []UserModel.User
	mysql.Db.Find(&userArr)

	for _, v := range userArr {
		if v.Id == 1 {
			data := UserModel.User{Nickname: "test", Remark: "remark", Birth: "123", Sex: "12", Height: 12,
				Hobby: "asdas"}
			userArr = append(userArr, data)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": userArr,
	})
}

func FirstUser(c *gin.Context) {
	user := new(UserModel.User)
	mysql.Db.First(user, 1)
	result := mysql.Db.Where("id = ?", 1).First(user).Value
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})
}

func SetRedis(c *gin.Context) {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()
	//redis.Redis.Set(ctx, "store", "12321", 0).Err()
	fmt.Println(redis.Redis.Get(ctx, "store").Val())
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}
