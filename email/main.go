package main

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"time"
)

// https://github.com/jordan-wright/email
func main() {

	ch := make(chan *email.Email, 100)
	p, err := email.NewPool(
		"smtp.gmail.com:587",
		4,
		smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"),
	)
	if err != nil {
		panic(err)
	}

	email := &email.Email{
		From:    fmt.Sprintf("发送者名字<%s>", "我是你爸爸"),
		To:      []string{"837098975.com@qq.com"},
		Subject: "这里是标题内容",
		Text:    []byte("这里是正文内容"),
	}
	ch <- email

	for i := 0; i < 4; i++ {
		go func() {
			for {
				select {
				case sendEmail := <-ch:
					p.Send(sendEmail, time.Millisecond)
				}
			}
		}()
	}
}
