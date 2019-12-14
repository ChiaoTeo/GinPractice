package userinfo

// Init 会初始化一个map，这个map保存着用户的信息
// 实际应用中，这部分数据会在数据库里保存
func Init() map[string]string {
	userinfo := make(map[string]string)
	userinfo["admin"] = "123456"
	return userinfo
}

// CheckValid 检查用户密码是否正确
func CheckValid(username string, password string) bool {
	info := Init()
	if info[username] == "" {
		return false
	}
	return info[username] == password
}
