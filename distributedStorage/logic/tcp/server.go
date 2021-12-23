package tcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"hb_distributeStorage/config"
	cache2 "hb_distributeStorage/logic/cache"
	"hb_distributeStorage/utils"
	"net"
)

func NewInitTcp() {
	listen, err := net.Listen("tcp", config.Config.Host+":"+config.Config.Port)
	if err != nil {
		fmt.Println("Listen() failed, err: ", err)
		return
	}
	for {
		conn, err := listen.Accept() // 监听客户端的连接请求
		if err != nil {
			fmt.Println("Accept() failed, err: ", err)
			continue
		}
		go process(conn) // 启动一个goroutine来处理客户端的连接请求
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err: ", err)
			break
		}
		recvStr := string(buf[:n])
		// 保存对应的文件加里面
		fmt.Println("收到Client端发来的数据：", recvStr)

		if string(buf[:9]) == "getMaster" {
			// 同步master节点信息
			aa, _ := cache2.LocalCache.Get("masterAddress")
			conn.Write([]byte(utils.Strval(aa)))
		} else if string(buf[:8]) == "NAddress" {
			// 同步节点信息
			syncNodeAddress(string(buf[8:]))
			aa, _ := cache2.LocalCache.Get("nodeAddress")
			mjson, _ := json.Marshal(utils.Strval(aa))
			conn.Write(mjson)
		} else {
			conn.Write([]byte("success"))
		}
		//else if string(buf[:8]) == "MAddress" {
		//	// 同步master节点信息
		//	syncMasterAddress(string(buf[:16]))
		//} else if string(buf[:8]) == "DAddress" {
		//	// 同步master节点信息
		//	syncDeleteAddress(string(buf[:16]))
		//}
	}
}

func IsTcpClient(address string) bool {
	conn, err := net.Dial("tcp4", address)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
