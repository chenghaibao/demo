package adminControler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"hb_gin/channel"
	mysql "hb_gin/plugin/mysql"
	redis "hb_gin/plugin/redis"
	"net/http"
	"strconv"
	sync2 "sync"

	UserModel "hb_gin/Model/User"
	AdminService "hb_gin/service/Admin"
)

type AdminController struct {
}

var mux sync2.Mutex
var wg sync2.WaitGroup

func (this *AdminController) GetPing(c *gin.Context) {
	sum := AdminService.GetPing(2)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "sadsa" + strconv.Itoa(sum),
	})
}

func (this *AdminController) AddUser(c *gin.Context) {
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

func (this *AdminController) DeleteUser(c *gin.Context) {
	user := new(UserModel.User)
	mysql.Db.Where("id = ?", 1).Delete(user)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

func (this *AdminController) UpdateUser(c *gin.Context) {
	// 更改1
	defer mux.Unlock()
	mux.Lock()
	mysql.Db.Model(&UserModel.User{}).Where("id = ?", 2).Update("nickname", "test")
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

func (this *AdminController) SelectUser(c *gin.Context) {
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

func (this *AdminController) FirstUser(c *gin.Context) {
	user := new(UserModel.User)
	mysql.Db.First(user, 1)
	result := mysql.Db.Where("id = ?", 1).First(user).Value
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": result,
	})
}

func (this *AdminController) SetRedis(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//redis.Redis.Set(ctx, "store", "12321", 0).Err()
	fmt.Println(redis.Redis.Get(ctx, "store").Val())
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

func (this *AdminController) SendChannel(c *gin.Context) {
	channel.AdminMessage <- "chenghaibao"
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

func (this *AdminController) SetMap(c *gin.Context) {
	// map  等于 struct  chan  interface 等等
	// chan 等于 struct  string  chan interface等等
	// 切片  等于 struct  string  chan interface等等
	hashMap := make([]map[string]string, 0)
	setMap := make(map[string]string)
	setMap["test"] = "test_ceshi"
	setMap["test1"] = "test1_ceshi"
	setMap["test2"] = "test2_ceshi"
	hashMap = append(hashMap, setMap)
	setMap1 := make(map[string]string)
	setMap1["test"] = "test_ceshi"
	setMap1["test1"] = "test1_ceshi"
	setMap1["test2"] = "test2_ceshi"
	hashMap = append(hashMap, setMap1)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": hashMap,
	})
}

func (this *AdminController) SetWaitGroup(c *gin.Context) {
	sum := 0
	for i := 1; i < 100; i++ {
		// 计数加 1
		wg.Add(1)
		go func(i int) {
			sum += i
			// 计数减 1
			defer wg.Done()
			fmt.Printf("goroutine%d 结束\n", i)
		}(i)
	}

	// 等待执行结束
	wg.Wait()
	fmt.Println("所有 goroutine 执行结束")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}

//定义interface
type VowelsFinder interface {
	FindVowels()
}

// strcut 可以定义 string 切片 chan map interface 啥都行
type MyString string

//实现接口
func (ms *MyString) FindVowels() {
	fmt.Println("sda")
}

func (this *AdminController) SetStruct(c *gin.Context) {
	name := MyString("Sam Anderson") // 类型转换
	var v VowelsFinder               // 定义一个接口类型的变量
	v = &name
	v.FindVowels()
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": true,
	})
}
