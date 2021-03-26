package Model

import (
	_struct "1/struct"
)

func UserInfo(uid string) (userInfo _struct.UserInfo,err error) {
	err = DB.Where("uid=?", uid).Find(&userInfo).Error
	//var userinfo _struct.UserInfo
	//row,err:=DB.Query("select * from userinfo where uid=?",uid);if err!=nil{
	//	return userinfo,err
	//}
	//defer row.Close()
	//
	//
	//for row.Next() {
	//	err = row.Scan(&userinfo.Uid,&userinfo.Stars,&userinfo.Fans,&userinfo.Likes,&userinfo.Plays,&userinfo.PictureURL)
	//	if err != nil {
	//		fmt.Printf("scan failed: %v", err)
	//		return userinfo,err
	//	}
	//}
	return userInfo,err
}
