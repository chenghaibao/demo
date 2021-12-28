package lib

import (
	"github.com/streadway/amqp"
	"log"
)

// RabbitMQ连接函数
func RabbitMQConn() (conn *amqp.Connection, err error) {
	// RabbitMQ分配的用户名称
	var user string = "rabbitmq"
	// RabbitMQ用户的密码
	var pwd string = "rabbitmq"
	// RabbitMQ Broker 的ip地址
	var host string = "127.0.0.1"
	// RabbitMQ Broker 监听的端口
	var port string = "5672"
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"
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
