package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"wechat/redis"
	"wechat/servers/websocket"
)



func SendMessage(c *gin.Context) {
	userId := c.Query("userId")
	data, _ := ioutil.ReadAll(c.Request.Body)

	client :=  websocket.ClientManagerins.GetUserClient(userId)
	client.SendMsg(string(data), userId)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
	})
}

func PullMessage(c *gin.Context) {
	userId := c.Query("userId")
	data, _ := redis.NewCache.Lrange("messageUserId_" + userId, 0 , -1)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
		"data": data,
	})
}

func Chat(c *gin.Context) {
	message := "你好客户端"
	client := websocket.ClientManagerins.GetUserClient("1001")
	if (client != nil) {
		client.Send <- []byte(message)
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
	})
}