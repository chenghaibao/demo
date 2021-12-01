package channel

import "hb_gin/common"

var AdminMessage chan string

func NewAdminMessageChannel(){
	AdminMessage = make(chan string,10)
	for{
		select {
			case adminMessage := <-AdminMessage:
				common.Log.Info(adminMessage)
		}
	}
}
