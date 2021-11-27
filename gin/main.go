package main

import (
	"hb_gin/config"
	"hb_gin/route"
)

func main(){
	config.NewConfig()
	gin := route.Routers()
	gin.Run(":3561")
}
