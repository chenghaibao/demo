package main

import (
	"hb_distributeStorage/config"
	"hb_distributeStorage/logic/cache"
	"hb_distributeStorage/logic/node"
	"hb_distributeStorage/logic/tcp"
	"hb_distributeStorage/work"
)

func main() {
	config.NewConfig()
	cache.NewCache()
	work.NewPool()
	node.AddNode()
	node.GetMaster()
	tcp.NewInitTcp()
	// 创建client 通道写入
}
