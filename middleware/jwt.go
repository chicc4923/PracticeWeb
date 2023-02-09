package jwt

import (
	"PracticeWeb/pkg/utils"
	"net/http"
	"time"

	"PracticeWeb/pkg/errors"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = errors.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = errors.INVALID_PARAMS
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = errors.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = errors.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != errors.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  errors.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
