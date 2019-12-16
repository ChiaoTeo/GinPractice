package users

import "GinPractice/conf"

import "fmt"

// CheckVaild 检查用户名与密码是否正确
func CheckVaild(username string, password string) (bool, error) {
	// 查询有没有usermae的用户
	rows, err := conf.DB.Query("select count(*) from users where username=?", username)
	if err != nil {
		fmt.Print(err)
		return false, err
	}
	var num interface{}
	// 遍历查出来的rows
	for rows.Next() {
		// 因为我们知道，count的结果只有一行一列，因此直接将结果存入num
		err = rows.Scan(&num)
		if err != nil {
			fmt.Print(err)
		}
		if num == nil {
			fmt.Print("缓存出错")
			return false, nil
		}
	}
	// 查出来的结果为0.说明不存在这个用户名的用户
	if num.(int64) == 0 {
		return false, nil
	}
	// 既然有用户。开始查询密码，并核对密码
	rows, err = conf.DB.Query("select * from users where username=?", username)
	col, err := rows.Columns()
	if err != nil {
		return false, err
	}
	// 这部分是一个普遍用法，如果有多列，就把多列都存下来 之所以只有
	// 这是每一列的储存数组
	res := make([]interface{}, len(col))
	//把每一列的储存数组的地址储存下来
	respo := make([]interface{}, len(col))
	for i := range res {
		respo[i] = &res[i]
	}
	// 每一列的数据储存下来
	for rows.Next() {
		rows.Scan(respo...)
	}
	// 将第二列（储存的密码）转化成string
	recodep := string(res[1].([]byte))
	// 核对密码
	if recodep == password {
		return true, nil
	}
	return false, nil
}
