package Model

import _struct "1/struct"

func SearchVideo(keyword string) (err error,videos []_struct.Video) {

	err = DB.Where("title like CONCAT('%',?,'%')", keyword).Find(&videos).Error

	//var sqlStr string
	//mp := make(map[int]_struct.Video)
	//
	//sqlStr ="select * from videos where title like CONCAT('%',?,'%');"
	//row,err:=DB.Query(sqlStr,keyword)
	//if err!=nil{
	//	return err,mp
	//}
	//defer row.Close()
	//
	//var v _struct.Video
	////最多8个视频
	//for i:=1;row.Next()&&i<=8;i++ {
	//	err = row.Scan(&v.ID,&v.UID,&v.Title,&v.Info,&v.Part,&v.VideoURL,&v.CoverURL,&v.UpTime)
	//	if err != nil {
	//		return err,mp
	//	}
	//	mp[i]=v
	//}
	return err,videos
}
