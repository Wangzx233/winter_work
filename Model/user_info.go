package Model

import _struct "1/struct"

func UInfo(uid string) (v _struct.UserInfo,err error) {


	v.PictureURL = "../public/imgs/default.jpg"//默认地址

	err = DB.Where("uid=?", uid).Find(&v).Error

	//sqlStr := "select * from user where uid=?"
	//row,err:=DB.Query(sqlStr,uid)
	//if err!=nil{
	//	return v,err
	//}
	//defer row.Close()

	//for row.Next() {
	//	err = row.Scan(&v.Uid,&v.Stars,&v.Fans,&v.Likes,&v.Plays,&v.PictureURL)
	//	if err != nil {
	//		return v,err
	//	}
	//}
	return v,err
}
