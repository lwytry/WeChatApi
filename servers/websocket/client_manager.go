package websocket

import (
	"fmt"
	"sync"
)

// 连接管理
type ClientManager struct {
	Clients     map[string]*Client // 登录的用户 // appId+uuid
	ClientsLock sync.RWMutex       // 读写锁
	Register    chan *Client       // 连接连接处理
	Unregister  chan *Client       // 断开连接处理程序
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
	}

	return
}


// 获取用户key
func GetUserKey(userId string) (key string) {
	key = fmt.Sprintf("%s", userId)

	return
}