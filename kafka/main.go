package main

import (
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	err := Init()
	if err != nil {
		panic(err)
	}
}

func Init() error {
	// kafka 是否链接成功
	broker := sarama.NewBroker("localhost:9092")
	err := broker.Open(nil)
	if err != nil {
		log.Println("error in open the broker, ", err)
		return err
	}

	err = broker.Close()
	if err != nil {
		log.Fatal("error in close admin, ", err)
		return err
	}
	return nil
}
