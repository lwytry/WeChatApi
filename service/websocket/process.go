package websocket

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Request struct {
	UserId  string      `json:"userId"`            // 用户一Id
	Cmd  	string      `json:"cmd"`            	// 请求命令字
	Data    interface{}
}

type DisposeFunc func(client *Client, seq string, message []byte) (code uint32, msg string, data interface{})

var (
	handlers        = make(map[string]DisposeFunc)
	handlersRWMutex sync.RWMutex
)

// 注册
func Register(key string, value DisposeFunc) {
	handlersRWMutex.Lock()
	defer handlersRWMutex.Unlock()
	handlers[key] = value

	return
}

func getHandlers(key string) (value DisposeFunc, ok bool) {
	handlersRWMutex.RLock()
	defer handlersRWMutex.RUnlock()

	value, ok = handlers[key]

	return
}

// 处理数据
func ProcessData(client *Client, message []byte) {

	fmt.Println("处理数据", client.Addr, string(message))

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("处理数据 stop", r)
		}
	}()

	request := &Request{}

	err := json.Unmarshal(message, request)
	if err != nil {
		fmt.Println("处理数据 json Unmarshal", err)
		//client.SendMsg([]byte("数据不合法"))
		return
	}

	requestData, err := json.Marshal(request.Data)
	if err != nil {
		fmt.Println("处理数据 json Marshal", err)
		//client.SendMsg([]byte("处理数据失败"))
		return
	}

	seq := request.UserId
	cmd := request.Cmd

	var (
		code uint32
		msg  string
		data interface{}
	)

	// 采用 map 注册的方式
	if value, ok := getHandlers(cmd); ok {
		code, msg, data = value(client, seq, requestData)
	} else {
		code = 1001
		fmt.Println("处理数据 路由不存在", client.Addr, "cmd", cmd)
	}

	fmt.Println("acc_response send", client.Addr, client.UserId, "cmd", cmd, "code", code, msg, data)

	return
}