package websocket

import (
	"fmt"
	"sync"
)

// 连接管理
type ClientManager struct {
	Clients     map[string]*Client // 登录的用户 // userId
	ClientsLock sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
}

// 管道处理程序
func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.Register:
			// 建立连接事件
			manager.EventRegister(conn)

		case conn := <-manager.Unregister:
			// 断开连接事件
			manager.EventUnregister(conn)
		}
	}
}

// 用户建立连接事件
func (manager *ClientManager) EventRegister(client *Client) {
	manager.AddClients(client)

	fmt.Println("EventRegister 用户建立连接", client.Addr)

	// client.Send <- []byte("连接成功")
}

// 用户断开连接
func (manager *ClientManager) EventUnregister(client *Client) {
	// 清除客户端
	manager.DelClients(client)

	fmt.Println("EventUnregister 用户断开连接", client.Addr, client.UserId)
}

// 删除客户端
func (manager *ClientManager) DelClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	delKey := client.UserId
	if _, ok := manager.Clients[delKey]; ok {
		delete(manager.Clients, delKey)
	}
}

func NewClientManager() (clientManager *ClientManager) {
	clientManager = &ClientManager{
		Clients:      make(map[string]*Client),
		Register:   make(chan *Client, 1000),
		Unregister: make(chan *Client, 1000),
	}

	return
}

// 添加客户端
func (manager *ClientManager) AddClients(client *Client) {
	manager.ClientsLock.Lock()
	defer manager.ClientsLock.Unlock()
	index := client.UserId;
	manager.Clients[index] = client
}

// 获取用户的连接
func (manager *ClientManager) GetUserClient(userId string) (client *Client) {

	manager.ClientsLock.RLock()
	defer manager.ClientsLock.RUnlock()

	userKey := GetUserKey(userId)
	if value, ok := manager.Clients[userKey]; ok {
		client = value
	} else {
		client = &Client{UserId:userId}
	}

	return
}


// 获取用户key
func GetUserKey(userId string) (key string) {
	key = fmt.Sprintf("%s", userId)

	return
}