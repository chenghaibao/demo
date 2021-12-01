package channel

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func TestChannel(t *testing.T) {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(true, true),
	)
	// 订阅当前topic
	messages, err := pubSub.Subscribe(context.Background(), "hb.topic")
	// 如果为错误则GG
	if err != nil {
		panic(err)
	}

	// 协程跑通道队列
	go process(messages)

	// 发送数据
	publishMessages(pubSub)
}

func publishMessages(publisher message.Publisher) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, world!"))

		if err := publisher.Publish("hb.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
		msg.Ack()
	}
}
