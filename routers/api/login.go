package api

import (
	"PracticeWeb/models"
	"PracticeWeb/pkg/errors"
	"PracticeWeb/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data models.User
	c.ShouldBindJSON(&data)
	var token string
	var code int
	isUser := models.CheckAuth(data.Username, data.Password)
	if isUser == true {
		token, _ = utils.GenerateToken(data.Username, data.Password)
		code = errors.SUCCESS
	} else {
		code = errors.ERROR_AUTH
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": data,
		"token":   token,
	})
}
