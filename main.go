package main

import (
	"fmt"
	"net/http"
	_ "wechat/model"
	"wechat/service/websocket"
	"wechat/router"
	"wechat/utils"
)

func main() {

	go websocket.StartWebSocket()
	httpRouter := router.InitRouter()
	router.WebsocketInit()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", utils.HTTPPort),
		Handler:        httpRouter,
		ReadTimeout:    utils.ReadTimeout,
		WriteTimeout:   utils.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}