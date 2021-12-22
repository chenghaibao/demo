package node

import (
	"github.com/patrickmn/go-cache"
	"hb_distributeStorage/config"
	cache2 "hb_distributeStorage/logic/cache"
	"hb_distributeStorage/logic/tcp"
	"hb_distributeStorage/utils"
	"strings"
	"sync"
)

var mux sync.Mutex

func getMaster() string {
	if address, ok := isExistNode(); ok {
		return address
	} else {
		return setMasterNode()
	}
}

func isExistNode() (string, bool) {
	address, ok := cache2.LocalCache.Get("masterAddress")
	// 请求master健康检查接口是否成功
	if true {
		return utils.Strval(address), ok
	} else {
		return "", false
	}

}

func setMasterNode() string {
	mux.Lock()
	defer mux.Unlock()
	// 地址config里面取
	array := []string{"127.0.0.1:9700", "127.0.0.1:9800", "127.0.0.1:9900"}
	address, _ := utils.Random(array, 1)

	// 通知删除某个节点
	array1 := utils.RemoveParam(array, utils.Strval(address))

	// 请求当前地址的健康接口
	//	curl.get
	// 不能接着剔除
	for _, v := range array1 {
		address, _ = utils.Random(array, 1)
		//	请求当前地址的健康接口

		//  可以返回
		address = v
		break
	}

	// 存入配置文件
	// 通知删除某个节点
	utils.RemoveParam(array, address)
	cache2.LocalCache.Set("masterAddress", address, cache.NoExpiration)

	// 同步选举数据
	synchronous()

	return address
}

func synchronous() {
	// 地址config里面取
	masterAddress, _ := cache2.LocalCache.Get("masterAddress")
	array := strings.Split(config.Config.Cluster, ",")
	synchronousArray := utils.RemoveParam(array, utils.Strval(masterAddress))
	for range synchronousArray {
		go func(value string) {
			// 同步node选举 到个节点 tcp链接传输
			tcp.SendClient(value)
		}(utils.Strval(masterAddress))
	}
}
