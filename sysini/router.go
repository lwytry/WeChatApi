package sysini

import (
	"github.com/gin-gonic/gin"
	"wechat/controller"
	"wechat/service"
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

	contact := apiv1.Group("/contact")
	contact.Use(service.JWT())
	{
		// 获取用户联系人
		contact.GET("/getList", controller.GetContactList)
	}

	chat := apiv1.Group("/chat")
	chat.Use(service.JWT())
	{
		// 获取聊天直播间token 用来发起视频通信
		chat.GET("/getRTCToken", controller.GetRTCToken);
		// 发送消息
		chat.POST("/message", controller.SendMessage)
		// 拉取消息
		chat.GET("/message", controller.PullMessage)
	}
	apiv1.GET("/getOSStoken", controller.GetSignUrl)

	return r
}