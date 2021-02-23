package Model

import (
	_struct "1/struct"
	"fmt"
)

func UserInfo(uid string) (_struct.UserInfo,error) {
	var userinfo _struct.UserInfo
	row,err:=DB.Query("select * from userinfo where uid=?",uid);if err!=nil{
		return userinfo,err
	}
	defer row.Close()


	for row.Next() {
		err = row.Scan(&userinfo.Uid,&userinfo.Stars,&userinfo.Fans,&userinfo.Likes,&userinfo.Plays,&userinfo.PictureURL)
		if err != nil {
			fmt.Printf("scan failed: %v", err)
			return userinfo,err
		}
	}
	return userinfo,err
}
