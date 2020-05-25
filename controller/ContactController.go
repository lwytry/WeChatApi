package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/model"
)

func GetContactList(c *gin.Context) {
	userId := c.Query("userId")
	if (userId == "") {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误",
			"errCode":1001,
		})
		return
	}
	ret := model.GetContactAll(userId)

	var retData  []model.User
	if (ret != nil) {
		for _, val := range ret  {
			retData = append(retData, val.User)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
		"data": retData,
	})
}