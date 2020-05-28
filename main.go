package main

import (
	"fmt"
	"net/http"
	_ "wechat/model"
	"wechat/service/websocket"
	"wechat/sysini"
	"wechat/utils"
)

func main() {

	go websocket.StartWebSocket()
	router := sysini.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", utils.HTTPPort),
		Handler:        router,
		ReadTimeout:    utils.ReadTimeout,
		WriteTimeout:   utils.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}