package common

import (
	"github.com/roylee0704/gron"
)

var cron *gron.Cron

func NewCron() *gron.Cron {
	cron = gron.New()
	//cron.AddFunc(gron.Every(1*time.Second), day.UpdateOrderStatus)
	return cron
}
