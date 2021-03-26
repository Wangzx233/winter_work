package Model

import (
	_struct "1/struct"
)

func Oauth(userinfo map[string]interface{},token string) bool {
	var user _struct.User
	DB.Where("token=?").Find(&user)
	//row,err:=DB.Query("select name from user where token=?",token)
	//defer row.Close()

	//for row.Next() {
	//	err = row.Scan(&u.UserName)
	//	if err != nil {
	//		fmt.Printf("scan failed: %v", err)
	//		return false
	//	}
	//}
	if user.Token==token {
		return true
	}else {
		DB.Create(&_struct.User{Password: "123456",Token: token})
		//_,err:=DB.Exec("insert user (name,password,token) values (?,?,?)",userinfo["id"],"123456",token)
		//if err!=nil{
		//	fmt.Printf("Exec failed: %v",err)
		//	return false
		//}
		return true
	}

}