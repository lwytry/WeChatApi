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
		apiv1.POST("/login", controller.Login)
		apiv1.GET("jwtParse", controller.ParseToken)
		apiv1.GET("/sendCaptcha", controller.SendCaptcha)
		apiv1.POST("/chat", controller.Chat)
	}

	return r
}