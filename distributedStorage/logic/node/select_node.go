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
	// 主节点选择子节点
	array := []string{"127.0.0.1:9700", "127.0.0.1:9800", "127.0.0.1:9900"}
	address, _ := utils.Random(array, 1)
	// 请求子节点（是否可以访问）
	ok := utils.IsTcpClient(address)
	if !ok {
		tcp.SendClient(address)
		newAddress := utils.RemoveParam(array, utils.Strval(address))
		for _, v := range newAddress {
			ok = utils.IsTcpClient(v)
			if ok {
				address = v
				break
			}
			// 不严谨 没有做所有节点都不存在的情况
		}
	}
	// 发送value
	fmt.Println(utils.GetCurl(utils.Strval(address) + "?aa=13"))
}

func AddNode() {
	// 查看节点是否已加入
	array := strings.Split(config.Config.Cluster, ",")
	// 判断所有子节点是否能访问
	// 访问更新cache
	nodeAddress := setNodeAddress()
	cache2.LocalCache.Set("nodeAddress", nodeAddress, cache.NoExpiration)
	// 发送自己节点信息
	for _, v := range array {
		tcp.SendClient(v)
		// 保存主节点
	}

}

//
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

func SyncNodeAddress(address string) {
	nodeAddress, _ := cache2.LocalCache.Get("nodeAddress")
	if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), address); !ok {
		nodeAddress.(map[string]interface{})[address] = address
	}
	cache2.LocalCache.Set("nodeAddress", nodeAddress.(map[string]interface{}), cache.NoExpiration)
}

func SyncDeleteAddress(address string) {
	nodeAddress, _ := cache2.LocalCache.Get("nodeAddress")
	if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), address); !ok {
		delete(nodeAddress.(map[string]interface{}), address)
	}
	cache2.LocalCache.Set("nodeAddress", nodeAddress.(map[string]interface{}), cache.NoExpiration)
}

func SyncMasterAddress(masterAddress string) {
	// 设置主节点信息
	cache2.LocalCache.Set("masterAddress", masterAddress, cache.NoExpiration)
	// 删除子节点
	nodeAddress, _ := cache2.LocalCache.Get("nodeAddress")
	if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), masterAddress); ok {
		delete(nodeAddress.(map[string]interface{}), masterAddress)
	}
	cache2.LocalCache.Set("nodeAddress", nodeAddress.(map[string]interface{}), cache.NoExpiration)
}
