package node

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"hb_distributeStorage/config"
	cache2 "hb_distributeStorage/logic/cache"
	"hb_distributeStorage/logic/tcp"
	"hb_distributeStorage/utils"
	"strings"
)

func SelectNode() {
	var address string
	// 主节点选择子节点
	if config.Config.Cluster == "" {
		address = config.Config.Host + ":" + config.Config.Port
	} else {
		array := strings.Split(config.Config.Cluster, ",")
		address, _ := utils.Random(array, 1)
		// 请求子节点（是否可以访问）
		ok := tcp.IsTcpClient(address)
		if !ok {
			tcp.SendClient(address, "DAddress")
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
	}
	// 发送value
	fmt.Println(utils.GetCurl(utils.Strval(address) + "?aa=13"))
}

func AddNode() {
	// 查看节点是否已加入
	array := strings.Split(config.Config.Cluster, ",")
	// 访问更新cache
	nodeAddress := setNodeAddress()
	cache2.LocalCache.Set("nodeAddress", nodeAddress, cache.NoExpiration)
	// 发送自己节点信息
	if array != nil {
		fmt.Println(nodeAddress, "-----加入节点")
		for _, v := range array {
			// 保存主节点
			tcp.SendClient(v, "NAddress")
		}
	}
}

func setNodeAddress() map[string]interface{} {
	nodeMap := make(map[string]interface{})
	array := strings.Split(config.Config.Cluster, ",")
	if nodeAddress, ok := cache2.LocalCache.Get("nodeAddress"); ok {
		for _, v := range array {
			if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), v); !ok {
				nodeAddress.(map[string]interface{})[v] = v
			}
		}
		return nodeAddress.(map[string]interface{})
	} else {
		for _, v := range array {
			nodeMap[v] = v
		}
		return nodeMap
	}
}
