package api

import (
	"PracticeWeb/models"
	"PracticeWeb/pkg/errors"
	"PracticeWeb/pkg/utils"
	"github.com/astaxie/beego/validation"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var test int

type auth struct {
	Username string `valid:"Required; MaxSize(20)"`
	Password string `valid:"Required; MaxSize(40)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := errors.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if test != 1 {
			code = errors.ERROR
		}
		if isExist {
			token, err := utils.GenerateToken(username, password)
			if err != nil {
				code = errors.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token

				code = errors.SUCCESS
			}

		} else {
			code = errors.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": test,
		"msg":  errors.GetMsg(code),
		"data": data,
	})
}
