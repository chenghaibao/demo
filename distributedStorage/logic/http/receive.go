package http

import "net/http"

func GetReceive(w http.ResponseWriter, r *http.Request) {
	// 读取信息并返回

	// 节点换算  在哪个目录下面
	w.Write([]byte("result code"))
}

func SetReceive(w http.ResponseWriter, r *http.Request) {
	// 节点换算  存在哪个目录下面

	// 同步数据至所有节点副本
	w.Write([]byte("result code"))
}

func deleteReceive(w http.ResponseWriter, r *http.Request) {
	// 删除信息并返回

	// 节点换算  在哪个目录下面
	w.Write([]byte("result code"))
}
