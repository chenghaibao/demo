package lib

import (
	"github.com/streadway/amqp"
	"log"
)

// RabbitMQ连接函数
func RabbitMQConn() (conn *amqp.Connection, err error) {
	// RabbitMQ分配的用户名称
	var user string = "MjphbXFwLWNuLTdtejJpNWtqajAwNTpMVEFJNXRNY0hENVNHbWR6Mzl3WE1MMWE="
	// RabbitMQ用户的密码
	var pwd string = "NjUyMTNCNTIwMDgxNTM0RDc2MUNGNzMwMEUyNUEwMzgwNENCRkYyMDoxNjQwOTMzMDA0OTc0"

	url := "amqp://" + user + ":" + pwd + "@" + "amqp-cn-7mz2i5kjj005.mq-amqp.cn-hangzhou-249959-a.aliyuncs.com" + "/test"
	// 新建一个连接
	conn, err = amqp.Dial(url)
	// 返回连接和错误
	return
}

// 错误处理函数
func ErrorHanding(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
