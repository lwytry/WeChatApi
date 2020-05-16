package sysini

import (
	"github.com/gin-gonic/gin"
	"wechat/controller"
	"wechat/servers"
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
		// 登录
		apiv1.POST("/login", controller.Login)
		// 发送验证码
		apiv1.GET("/sendCaptcha", controller.SendCaptcha)
		// 更新token时间
		apiv1.GET("/refreshToken", controller.RefreshToken)
	}

	chat := apiv1.Group("/chat")
	chat.Use(servers.JWT())
	{
		// 发送消息
		chat.POST("/message", controller.SendMessage)
		// 拉取消息
		chat.GET("/message", controller.PullMessage)
	}

	return r
}