package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"wechat/model"
	"wechat/redis"
)

func GetUserlist(c *gin.Context) {
	//list := make(map[string]interface{})
	list := model.GetUserList(1, 10, nil)
	c.JSON(http.StatusOK, gin.H{
		"message": c.Query("username"),
		"code": http.StatusOK,
		"data": list,
	})

}

func Register(c *gin.Context) {
	phoneNum := c.PostForm("phone")
	codeString := c.PostForm("code");
	code, _ := strconv.Atoi(codeString)
	fmt.Print(code)
	ret := model.AddUser(phoneNum, int(time.Now().Unix()))
	if (ret) {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
			"errCode": 0,
		})
	}
}

func redisTest() {
	_, err := redis.NewCache.SetString("aaa", 33)
	fmt.Println(err)
	val, _ := redis.NewCache.GetInt("aaa")
	fmt.Println(val)
}



