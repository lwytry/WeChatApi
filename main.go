package main

import (
	"fmt"
	"net/http"
	"wechat/sysini"
	_ "wechat/model"
	"wechat/utils"
)

func main() {
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