package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
type HeartBeat struct {
	UserId string `json:"userId,omitempty"`
}
// 心跳接口
func HeartbeatController(client *Client, seq string, message []byte) (code uint32, msg string, data interface{}) {

	code = http.StatusOK
	currentTime := uint64(time.Now().Unix())

	request := &HeartBeat{}
	if err := json.Unmarshal(message, request); err != nil {
		code = 1001
		fmt.Println("心跳接口 解析数据失败", seq, err)

		return
	}

	fmt.Println("webSocket_request 心跳接口", client.UserId, client.Addr)

	client.Heartbeat(currentTime)

	return
}