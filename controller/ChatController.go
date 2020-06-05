package controller

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
	"wechat/redis"
	"wechat/service/rtc"
	"wechat/service/websocket"
)



func SendMessage(c *gin.Context) {
	dstId := c.Query("dstId")
	data, _ := ioutil.ReadAll(c.Request.Body)

	client :=  websocket.ClientManagerins.GetUserClient(dstId)
	client.SendMsg(string(data), dstId)
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

func GetRTCToken(c *gin.Context) {
	userId := c.Query("userId")
	roomId := c.Query("roomId")
	if (userId == "" || roomId == "") {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误",
			"errCode": 1001,
		})
		return
	}
	expireTime := time.Now().Unix() + 3600 * 2
	manager := rtc.NewManager()
	access := rtc.RoomAccess{
		RoomName: roomId,
		UserID: userId,
		ExpireAt: expireTime,
		Permission: "user",
	}
	token, err := manager.GetRoomToken(access)
	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": "生成失败",
			"errCode": 1001,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
		"data": token,
	})
}