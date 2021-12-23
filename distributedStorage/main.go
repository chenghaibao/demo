package main

import (
	"hb_distributeStorage/config"
	"hb_distributeStorage/logic/cache"
	"hb_distributeStorage/logic/node"
	"hb_distributeStorage/logic/tcp"
)

func main() {
	config.NewConfig()
	cache.NewCache()
	node.AddNode()
	node.GetMaster()
	tcp.NewInitTcp()
}
