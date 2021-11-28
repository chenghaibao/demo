package UserModel

import (
	"fmt"
	"hb_gin/initialize/mysql"
)

type User struct {
	Id       int         `gorm:"primary_key" json:"id"`
	Nickname string      `json:"nickname,omitempty"`
	Birth    string      `json:"birth"`
	Sex      string      `json:"sex"`
	Height   int         `json:"height"`
	Hobby    string      `json:"hobby"`
	Remark   interface{} `json:"remark"`
}

//添加数据
func (User *User) Add() {
	err := mysql.Db.Create(User).Error
	if err != nil {
		fmt.Println("创建失败")
	}
}
