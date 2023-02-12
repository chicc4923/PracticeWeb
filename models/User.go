package models

type User struct {
	Userid   int    `gorm:"primary_key" json:"userid"`
	Username string `gorm:"username" json:"username" banding:"require"`
	Password string `gorm:"password" json:"password" banding:"require"`
	Email    string `gorm:"email" json:"email"`
}

// ExistUser 查询用户是否存在
func ExistUser(email string) bool {
	var user User
	//db.Select("userid").Where("username=?", username).First(&user)
	db.Where("email=?", email).First(&user)
	if user.Userid > 0 {
		return true
	}
	return false
}

// AddUser 新增用户
func AddUser(user *User) bool {
	err := db.Create(&user)
	if err != nil {
		return false
	}

	return true
}

func CheckAuth(username, password string) bool {
	var user User
	db.Select("id").Where(User{Username: username, Password: password}).First(&user)
	if user.Userid > 0 {
		return true
	}

	return false
}
