package node

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"hb_distributeStorage/config"
	cache2 "hb_distributeStorage/logic/cache"
	"hb_distributeStorage/logic/tcp"
	"hb_distributeStorage/utils"
	"hb_distributeStorage/work"
	"strings"
	"sync"
)

var mux sync.Mutex

// 通过
func GetMaster() string {
	if address, ok := isExistNode(); ok {
		fmt.Println("getMaster", address)
		return address
	} else {
		return setMasterNode()
	}
}

// 通过
func isExistNode() (string, bool) {
	address, ok := cache2.LocalCache.Get("masterAddress")
	// 请求master健康检查接口是否成功
	if ok {
		return utils.Strval(address), ok
	} else {
		// 获取远程节点的master
		array := strings.Split(config.Config.Cluster, ",")
		for _, v := range array {
			masterAddress := tcp.SendClient(v, "getMaster")
			if masterAddress != "success" {
				cache2.LocalCache.Set("masterAddress", masterAddress, cache.NoExpiration)
				return utils.Strval(address), ok
			}
			// 不严谨 没有做所有节点都不存在的情况
		}
		return "", false
	}
}

// 通过
func setMasterNode() string {
	mux.Lock()
	defer mux.Unlock()
	// 地址config里面取
	array := strings.Split(config.Config.Cluster, ",")
	address, _ := utils.Random(array, 1)
	if address == "" {
		cache2.LocalCache.Set("masterAddress", config.Config.Host+":"+config.Config.Port, cache.NoExpiration)
	} else {
		// 请求当前地址的健康接口
		ok := tcp.IsTcpClient(address)
		if !ok {
			work.Pool.Pool <- address + "DAddress"
			newAddress := utils.RemoveParam(array, utils.Strval(address))
			for _, v := range newAddress {
				ok = tcp.IsTcpClient(v)
				if ok {
					address = v
					break
				}
				// 不严谨 没有做所有节点都不存在的情况
			}
		}
		// 通知删除某个节点
		utils.RemoveParam(array, address)
		cache2.LocalCache.Set("masterAddress", address, cache.NoExpiration)
	}
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
			tcp.SendClient(value, "MAddress")
		}(utils.Strval(masterAddress))
	}
}
