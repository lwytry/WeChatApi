package router

import (
	"wechat/service/websocket"
)

func WebsocketInit() {
	websocket.Register("heartbeat", websocket.HeartbeatController)
	//websocket.Register("ping", websocket.PingController)
}
