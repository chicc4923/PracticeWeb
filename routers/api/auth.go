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

type User struct {
	Userid   int    `gorm:"primary_key" json:"userid"`
	Username string `gorm:"username" json:"username" banding:"require"`
	Password string `gorm:"password" json:"password" banding:"require"`
	Email    string `gorm:"email" json:"email"`
}

func GetAuth(c *gin.Context) {
	var user User
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	//a := User{Username: username, Password: password}
	user.Username = username
	user.Password = password

	ok, _ := valid.Valid(&user)

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
