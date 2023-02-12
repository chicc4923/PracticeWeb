package api

import (
	"PracticeWeb/models"
	"PracticeWeb/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddUser Add new user to mysql
func AddUser(c *gin.Context) {
	var data models.User
	_ = c.ShouldBindJSON(&data)
	code := errors.INVALID_PARAMS
	exist := models.ExistUser(data.Email)

	// if email exists,return error
	if exist == true {
		code = errors.ERROR_EXIST_Email
	} else {
		code = errors.SUCCESS
		models.AddUser(&data)
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"data":    data,
		"message": errors.GetMsg(code),
	})
}
