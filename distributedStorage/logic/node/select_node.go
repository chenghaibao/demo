package node

import (
	"fmt"
	"hb_distributeStorage/utils"
)

func SelectNode() {
	// 主节点选择子节点
	array := []string{"127.0.0.1:9700", "127.0.0.1:9800", "127.0.0.1:9900"}
	address, _ := utils.Random(array, 1)
	// 请求子节点（是否可以访问）

	// 发送value

	//curl address
	fmt.Println(address)
}

func addNode() {
	// 查看节点是否已加入
	array := []string{"127.0.0.1:9700", "127.0.0.1:9800", "127.0.0.1:9900"}

	// 没加入则加入节点   // 否则不处理
	address, _ := utils.Random(array, 1)
	// 请求子节点（是否可以访问）

	// 发送value

	//curl address
	fmt.Println(address)
}
