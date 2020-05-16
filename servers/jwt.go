package servers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wechat/utils"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		j := &utils.JWT{[]byte(utils.SignKey)}
		code = 0
		token := c.GetHeader("token");
		if token == "" {
			code = 1001
		} else {
			_, err := j.ParseToken(token)
			if err != nil {
				code = 1002
			}
		}

		if code != 0 {
			c.JSON(http.StatusOK, gin.H{
				"message": "",
				"errCode": code,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}