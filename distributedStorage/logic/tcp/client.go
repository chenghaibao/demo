package tcp

import (
	"fmt"
	"hb_distributeStorage/logic/node"
	"net"
	"strings"
)

func SendClient(input string){
	conn, err := net.Dial("tcp4", "127.0.0.1:9997")
	if err != nil {
		fmt.Println("err : ", err)
		return
	}
	defer conn.Close() // 关闭TCP连接
	for {
		inputInfo := strings.Trim(input, "\r\n")
		_, err := conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [4096]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		if string(buf[:8]) == "NAddress"{
			// 同步节点信息
			node.SyncNodeAddress(string(buf[:16]))
		}

		if string(buf[:8]) == "MAddress"{
			// 同步master节点信息
			node.SyncMasterAddress(string(buf[:16]))
		}

		if string(buf[:8]) == "DAddress"{
			// 同步master节点信息
			node.SyncDeleteAddress(string(buf[:16]))
		}

		fmt.Println(string(buf[:n]))
	}
}