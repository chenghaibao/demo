package tcp

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	cache2 "hb_distributeStorage/logic/cache"
	"hb_distributeStorage/utils"
	"net"
	"strings"
)

func SendClient(address string, input string) string {
	conn, err := net.Dial("tcp4", address)
	if err != nil {
		fmt.Println("err : ", err)
		return ""
	}
	//for {
	inputInfo := strings.Trim(input+address, "\r\n")
	_, err = conn.Write([]byte(inputInfo)) // 发送数据
	if err != nil {
		conn.Close() // 关闭TCP连接
		return ""
	}
	buf := [4096]byte{}
	n, err := conn.Read(buf[:])
	if err != nil {
		conn.Close()
		fmt.Println("recv failed, err:", err)
		return ""
	}
	//if string(buf[:8]) == "NAddress" {
	//	// 同步节点信息
	//	syncNodeAddress(string(buf[:16]))
	//}
	//
	//if string(buf[:8]) == "MAddress" {
	//	// 同步master节点信息
	//	syncMasterAddress(string(buf[:16]))
	//}
	//
	//if string(buf[:8]) == "DAddress" {
	//	// 同步master节点信息
	//	syncDeleteAddress(string(buf[:16]))
	//}
	//
	//if string(buf[:8]) == "getMaster" {
	//	// 同步master节点信息
	//	syncGetAddress(string(buf[:16]))
	//}
	fmt.Println(string(buf[:n]))
	return string(buf[:n])

}

func syncNodeAddress(address string) {
	nodeAddress, _ := cache2.LocalCache.Get("nodeAddress")
	if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), address); !ok {
		nodeAddress.(map[string]interface{})[address] = address
	}
	fmt.Println(nodeAddress, "1121321")
	cache2.LocalCache.Set("nodeAddress", nodeAddress.(map[string]interface{}), cache.NoExpiration)
}

func syncDeleteAddress(address string) {
	nodeAddress, _ := cache2.LocalCache.Get("nodeAddress")
	if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), address); !ok {
		delete(nodeAddress.(map[string]interface{}), address)
	}
	cache2.LocalCache.Set("nodeAddress", nodeAddress.(map[string]interface{}), cache.NoExpiration)
}

func syncMasterAddress(masterAddress string) {
	// 设置主节点信息
	cache2.LocalCache.Set("masterAddress", masterAddress, cache.NoExpiration)
	// 删除子节点
	nodeAddress, _ := cache2.LocalCache.Get("nodeAddress")
	if ok, _ := utils.MapKeyExist(nodeAddress.(map[string]interface{}), masterAddress); ok {
		delete(nodeAddress.(map[string]interface{}), masterAddress)
	}
	cache2.LocalCache.Set("nodeAddress", nodeAddress.(map[string]interface{}), cache.NoExpiration)
}

func syncGetAddress(masterAddress string) {
	// 设置主节点信息
	cache2.LocalCache.Set("masterAddress", masterAddress, cache.NoExpiration)
}
