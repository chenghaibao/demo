package main

import (
	"hb_distributeStorage/logic/cache"
	"hb_distributeStorage/logic/node"
)

func main() {
	cache.NewCache()
	node.AddNode()
}
