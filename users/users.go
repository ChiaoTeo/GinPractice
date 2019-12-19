package users

import (
	"GinPractice/model"
)

// CheckVaild 检查用户名与密码是否正确
func CheckVaild(username string, password string) (bool, error) {
	// 查询有没有usermae的用户
	var user model.User
	if err := model.DB.Where("username=?", username).First(&user).Error; err != nil {
		return false, err
	}
	if user.Password == password {
		return true, nil
	}
	return false, nil
}
