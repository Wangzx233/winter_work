package Model

import (
	_struct "1/struct"
)

func Oauth(userinfo map[string]interface{},token string) (bool,uint) {
	var user _struct.User
	DB.Where("token=?",token).Find(&user)
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
		return true,user.ID
	}else {

		DB.Create(&_struct.User{Password: "123456",Token: token,UserInfo: _struct.UserInfo{
			Stars:      0,
			Fans:       0,
			Likes:      0,
			Plays:      0,
			PictureURL: "",
		}})
		DB.Where("token=?",token).Find(&user)
		//_,err:=DB.Exec("insert user (name,password,token) values (?,?,?)",userinfo["id"],"123456",token)
		//if err!=nil{
		//	fmt.Printf("Exec failed: %v",err)
		//	return false
		//}
		return true,user.ID
	}

}