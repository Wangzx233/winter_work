package Model

import _struct "1/struct"

func Register(username string, password string) (bool, string) {
	err := DB.Where("user_name = ?", username).First(&_struct.User{}).Error
	if err == nil {
		return false, "用户已存在"
	}
	err = DB.Create(&_struct.User{UserName: username, Password: password, UserInfo: _struct.UserInfo{Uid: username}}).Error
	if err!=nil {
		return true, "创建成功"
	}
	return false,"创建失败"
}
func Login(username string, password string) bool {
	err := DB.Where("user_name = ? and password = ?", username, password).First(&_struct.User{}).Error
	if err == nil {
		return true
	}
	return false
}

//func Login(username string,password string) bool {
//	sqlStr:="select name,password from user where name=?"
//	row,err:=DB.Query(sqlStr,username)
//	if err!=nil{
//		fmt.Println("false")
//		return false
//	}
//	defer row.Close()
//	var u _struct.User
//	for row.Next() {
//		err = row.Scan(&u.UserName, &u.Password)
//		if err != nil {
//			fmt.Printf("scan failed: %v", err)
//			return false
//		}
//	}
//	if u.Password == password {
//		return true
//	}
//	return false
//}
//func Register(username string,password string) (bool,string) {
//	row,err:=DB.Query("select name from user where name=?",username)
//	defer row.Close()
//	var u _struct.User
//	for row.Next() {
//		err = row.Scan(&u.UserName)
//		if err != nil {
//			fmt.Printf("scan failed: %v", err)
//			return false,"查找失败"
//		}
//	}
//	if u.UserName==username {
//		return false,"id已被占用"
//	}
//	if !R_update(username,password){
//		return false,"创建失败"
//	}
//	return true,"创建成功"
//}
////新建用户
//func R_update(username,password string) bool {
//
//	_,err:=DB.Exec("insert user (name,password) values (?,?)",username,password)
//	if err!=nil{
//		fmt.Printf("Exec failed: %v",err)
//		return false
//	}
//	return true
//}
