package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	ClientManagerins = NewClientManager() // 管理者
	appIds        = []uint32{101, 102} // 全部的平台

	serverIp   string
	serverPort string
)

// 启动程序
func StartWebSocket() {

	http.HandleFunc("/acc", wsPage)

	// 添加处理程序
	go ClientManagerins.start()

	http.ListenAndServe(":8001", nil)
}

func wsPage(w http.ResponseWriter, req *http.Request) {

	// 升级协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])

		return true
	}}).Upgrade(w, req, nil)
	if err != nil {
		http.NotFound(w, req)

		return
	}

	fmt.Println("webSocket 建立连接:", conn.RemoteAddr().String())
	// 连接处可以做认证
	userId := req.URL.Query().Get("userId");

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, userId, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	ClientManagerins.Register <- client
}

