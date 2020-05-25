package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"wechat/model"
	"wechat/redis"
	"wechat/utils"
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
	code := c.PostForm("code");
	if (phoneNum == "" || code == "") {
		c.JSON(http.StatusOK, gin.H{
			"message": "参数错误",
			"errCode":1001,
		})
		return
	}
	// 用于验证码校验
	redisCoder, _ := redis.NewCache.GetBytes(phoneNum)
	if (code != string(redisCoder)) {
		c.JSON(http.StatusOK, gin.H{
			"message": "验证码错误",
			"errCode":1002,
		})
		return
	}

	uid := uuid.NewV1().String()
	encodeString := base64.StdEncoding.EncodeToString([]byte(uid))
	identifier := "wx_" + encodeString[:17]
	ret := model.AddUser(phoneNum, identifier, int(time.Now().Unix()))
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

func SendCaptcha(c *gin.Context) {
	phoneNum := c.Query("phone")
	fmt.Println(phoneNum)
	if (phoneNum == "" ) {
		c.JSON(http.StatusOK, gin.H{
			"message": "手机号错误",
			"errCode":1001,
		})
		return
	}
	// 发送验证码


	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	_, err := redis.NewCache.SetString(phoneNum, vcode)
	if (err != nil) {
		c.JSON(http.StatusOK, gin.H{
			"message": "发送失败",
			"errCode":1002,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "发送成功",
		"errCode":0,
	})
}



func Login(c *gin.Context) {
	phoneNum := c.PostForm("phone")
	codeString := c.PostForm("code");
	code, _ := strconv.Atoi(codeString)
	user := model.GetUserByPhone(phoneNum)
	if (user.ID == 0) {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户账号错误",
			"errCode":1001,
		})
		return
	}
	coder, _ := redis.NewCache.GetInt(phoneNum)
	if (coder != code) {
		c.JSON(http.StatusOK, gin.H{
			"message": "验证码错误",
			"errCode":1002,
		})
		return
	}

	//redis.NewCache.DelKey(phoneNum)
	j := &utils.JWT{[]byte(utils.SignKey)}
	claims := utils.UserInfo{
		strconv.Itoa(int(user.ID)), user.Identifier, user.Name, user.Phone,jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600*24),
			Issuer:    "lwy",
		},
	}

	token, _ := j.CreateToken(claims)
	data := make(map[string] interface{})
	data["token"] = token
	data["userInfo"] = user
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode":0,
		"data": data,
	})
}
func RefreshToken(c *gin.Context) {
	j := &utils.JWT{[]byte(utils.SignKey)}
	token := c.GetHeader("token")
	newToken, err := j.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "生成token失败",
			"errCode":1001,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"errCode": 0,
		"data": newToken,
	})
}

