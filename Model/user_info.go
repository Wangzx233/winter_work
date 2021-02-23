package Model

import _struct "1/struct"

func UInfo(uid string) (_struct.UserInfo,error) {
	var v _struct.UserInfo

	v.PictureURL = "https://i1.hdslb.com/bfs/face/b9c4f94842401e17419689789d5e02f6260da9d4.jpg@140w_140h_1c.webp"
	sqlStr := "select * from user where uid=?"
	row,err:=DB.Query(sqlStr,uid)
	if err!=nil{
		return v,err
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&v.Uid,&v.Stars,&v.Fans,&v.Likes,&v.Plays,&v.PictureURL)
		if err != nil {
			return v,err
		}
	}
	return v,err
}
