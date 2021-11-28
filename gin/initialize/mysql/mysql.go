package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hb_gin/config"
)

var Db *gorm.DB

func NewMysql() {
	// "root:u4bVDgdvELq6Nmhb@tcp(192.168.188.12:33061)/project-family"
	// root123456Hb@@
	mysqlConfig := config.Conf.Mysql.UserName + ":" + config.Conf.Mysql.Password + "@tcp(" + config.Conf.Mysql.Path + ":" + config.Conf.Mysql.Port + ")/" + config.Conf.Mysql.DbName
	fmt.Println(mysqlConfig)
	db, err := gorm.Open("mysql", mysqlConfig)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "hb_" + defaultTableName
	}
	Db = db
}
