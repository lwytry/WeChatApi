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

	userId := req.URL.Query().Get("userId");

	currentTime := uint64(time.Now().Unix())
	client := NewClient(conn.RemoteAddr().String(), conn, userId, currentTime)

	go client.read()
	go client.write()

	// 用户连接事件
	ClientManagerins.Register <- client
}

// 管道处理程序
func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)

		//case login := <-manager.Login:
		//	// 用户登录
		//	manager.EventLogin(login)
		//
		//case conn := <-manager.Unregister:
		//	// 断开连接事件
		//	manager.EventUnregister(conn)
		//
		//case message := <-manager.Broadcast:
		//	// 广播事件
		//	clients := manager.GetClients()
		//	for conn := range clients {
		//		select {
		//		case conn.Send <- message:
		//		default:
		//			close(conn.Send)
		//		}
		//	}
		}
	}
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)

	fmt.Println("EventRegister 用户建立连接", client.Addr)

	// client.Send <- []byte("连接成功")
}