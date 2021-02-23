package Model

import (
	_struct "1/struct"
)
//主页视频
func FindVideo(part string) (error,map[int]_struct.Video) {
	var sqlStr string
	mp := make(map[int]_struct.Video)

	sqlStr ="select * from videos where part=?"
	row,err:=DB.Query(sqlStr,part)
	if err!=nil{
		return err,mp
	}
	defer row.Close()

	var v _struct.Video
	//最多8个视频
	for i:=1;row.Next()&&i<=8;i++ {
		err = row.Scan(&v.ID,&v.UID,&v.Title,&v.Info,&v.Part,&v.VideoURL,&v.CoverURL,&v.UpTime)
		if err != nil {
			return err,mp
		}
		mp[i]=v
	}
	return err,mp
}