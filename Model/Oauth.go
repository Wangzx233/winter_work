package Model

import (
	_struct "1/struct"
	"fmt"
)

func Oauth(userinfo map[string]interface{},token string) bool {
	row,err:=DB.Query("select name from user where token=?",token)
	defer row.Close()
	var u _struct.User
	for row.Next() {
		err = row.Scan(&u.UserName)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return false
		}
	}
	if u.Token==token {
		return true
	}else {
		_,err:=DB.Exec("insert user (name,password,token) values (?,?,?)",userinfo["id"],"123456",token)
		if err!=nil{
			fmt.Printf("Exec failed: %v",err)
			return false
		}
		return true
	}

}