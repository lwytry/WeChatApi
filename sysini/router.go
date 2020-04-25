package sysini

import (
	"github.com/gin-gonic/gin"
	"wechat/controller"
	"wechat/utils"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(utils.RunMode)
	apiv1 := r.Group("/v1")
	{
		// 注册
		apiv1.POST("/register", controller.Register)
		// 查询列表
		apiv1.GET("/getuserlist", controller.GetUserlist)
	}

	return r
}