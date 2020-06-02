package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/service/aliyun"
)

func GetSignUrl(c *gin.Context) {
	userId := c.Query("userId")
	fileType := c.Query("fileType")

	if (userId == "" || fileType == "") {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误",
			"errCode": 1001,
		})
		return
	}
	signed, err := aliyun.GetSignUrl(userId + "/" + fileType + ".jpeg", 60)
	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": "生成失败",
			"errCode": 1002,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
		"data": signed,
	})
}
