package http

import "net/http"

func NewInitHttp() {
	http.ListenAndServe("127.0.0.1:9700", nil)
	http.HandleFunc("/setRecive", testReceive)
	http.HandleFunc("/getRecive", testReceive)
	http.HandleFunc("/deleteRecive", testReceive)
	http.HandleFunc("/health", testReceive)
}

func testReceive(w http.ResponseWriter, r *http.Request) {

	// 选举方法

	// 存在则通过master 随机访问节点进行处理

	// 保存到副本里面
	// tcp 同步数据到各个节点去

	// 保存master

}
