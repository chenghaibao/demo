package channel

import (
	"hb_gin/common"
	"sync/atomic"
)

var AdminMessage chan string

func NewAdminMessageChannel() {
	AdminMessage = make(chan string, 10)
	for {
		select {
		case adminMessage := <-AdminMessage:
			atomic.AddInt32(&chanLength, 1)
			common.Log.Info(adminMessage)
		}
	}
}
