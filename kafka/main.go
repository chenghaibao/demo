package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"time"
)

func main() {
	err := Create("hb_test")
	if err != nil {
		panic(err)
	}
}

func Create(topicName string) error {
	log.Println("start create topic...")

	broker := sarama.NewBroker("localhost:9092")
	err := broker.Open(nil)
	if err != nil {
		log.Println("error in open the broker, ", err)
		return err
	}
	//
	//var topicDetail sarama.TopicDetail
	//config := sarama.NewConfig()
	////config.Version = sarama.V2_1_0_0  //kafka版本号
	//config.Net.SASL.Enable = true
	//config.Net.SASL.Mechanism = "PLAIN"
	//config.Net.SASL.User = "admin"
	//config.Net.SASL.Password = "admin"
	//
	//admin, err := sarama.NewClusterAdmin([]string{"localhost:9092"}, config)
	//if err != nil {
	//	log.Fatal("error in create new cluster ... ", err)
	//	return err
	//}
	//
	//err = admin.CreateTopic(topicName, &topicDetail, false)
	//if err != nil {
	//	log.Println("error in create topic, ", err)
	//	return err
	//}
	provide()
	err = broker.Close()
	if err != nil {
		log.Fatal("error in close admin, ", err)
		return err
	}
	return nil
}

func provide() {
	//配置发布者
	config := sarama.NewConfig()
	//确认返回，记得一定要写，因为本次例子我用的是同步发布者
	config.Producer.Return.Successes = true
	//设置超时时间 这个超时时间一旦过期，新的订阅者在这个超时时间后才创建的，就不能订阅到消息了
	config.Producer.Timeout = 5 * time.Second
	//连接发布者，并创建发布者实例
	p, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Printf("sarama.NewSyncProducer err, message=%s \n", err)
		return
	}
	//程序退出时释放资源
	defer p.Close()
	//设置一个逻辑上的分区名，叫安彦飞
	topic := "hb_test"
	//这个是发布的内容
	srcValue := "sync: this is a message. index=%d"
	//发布者循环发送0-9的消息内容
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf(srcValue, i)
		//创建发布者消息体
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(value),
		}
		//发送消息并返回消息所在的物理分区和偏移量
		partition, offset, err := p.SendMessage(msg)
		if err != nil {
			log.Printf("send message(%s) err=%s \n", value, err)
		} else {
			fmt.Fprintf(os.Stdout, value+"发送成功，partition=%d, offset=%d \n", partition, offset)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
