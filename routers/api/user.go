package api

import (
	"PracticeWeb/models"
	"PracticeWeb/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	exsit := models.ExistUser(data.Username)
	if !exsit {
		models.AddUser(&data)
	} else {
		//TODO:用户存在时前端提示用户名已存在
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "add user success",
		"data":    data,
		"message": errors.GetMsg(200),
	})
}
