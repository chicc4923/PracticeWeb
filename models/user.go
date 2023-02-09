package models

import "PracticeWeb/pkg/errors"

type User struct {
	UserID   int    `gorm:"primarykey" json:"userID"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `gorm:"email" json:"email"`
}

// ExistUser 查询用户是否存在
func ExistUser(name string) bool {
	var user User
	db.Select("userid").Where("username=?", name).First(&user)
	if user.UserID > 0 {
		return false
	}
	return true
}

// AddUser 新增用户
func AddUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

//TODO:用户路由仍有问题
