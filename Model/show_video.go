package Model

import (
	_struct "1/struct"
	"log"
)

func ShowVideo(id int) (video _struct.Video, err error) {

	err = DB.Where("id=?", id).Find(&video).Error

	//sqlStr := "select * from videos where id=?"
	//row,err:=DB.Query(sqlStr,id)
	//if err!=nil{
	//	return v,err
	//}
	//defer row.Close()
	//
	//for row.Next() {
	//	err = row.Scan(&v.ID,&v.UID,&v.Title,&v.Info,&v.Part,&v.VideoURL,&v.CoverURL,&v.UpTime)
	//	if err != nil {
	//		return v,err
	//	}
	//}
	return video, err
}

func VideoInfo(vid int) (videoInfo _struct.VideoInfo, err error) {
	err = DB.Where("vid=?", vid).Find(&videoInfo).Error
	//sqlStr := "select * from video_info where vid=?"
	//row,err:=DB.Query(sqlStr,vid)
	//if err!=nil{
	//	return v,err
	//}
	//defer row.Close()
	//
	//for row.Next() {
	//	err = row.Scan(&v.ID,&v.UID,&v.Plays,&v.Likes,&v.Coins,&v.Collections,&v.Online)
	//	if err != nil {
	//		return v,err
	//	}
	//}
	return videoInfo,err
}

func OnlineAdd(vid string) {
	vInfo :=_struct.VideoInfo{}
	err := DB.Where("vid=?", vid).Model(&_struct.VideoInfo{}).Update("online", "online+1").Update("plays", "plays+1").Find(&vInfo).Error
	if err!=nil {
		log.Panicln(err)
	}
	err = DB.Where("uid=?", vInfo.UID).Model(&_struct.UserInfo{}).Update("plays", "plays+1").Error
	if err!=nil {
		log.Panicln(err)
	}
	//result, err := DB.Exec("update video_info set online = online+1,plays=plays+1 where vid = ?", vid)
	//if err != nil {
	//	fmt.Printf("update faied, error:[%v]", err.Error())
	//	return
	//}
	//
	//uid, err := result.LastInsertId()
	//if err != nil {
	//	fmt.Printf("error:[%v]", err.Error())
	//	return
	//}
	//_, err = DB.Exec("update user_info set plays = plays+1 where uid = ?", uid)
	//if err != nil {
	//	fmt.Printf("update faied, error:[%v]", err.Error())
	//	return
	//}

}
func OnlineSub(vid string) {
	err := DB.Where("vid=?", vid).Model(&_struct.VideoInfo{}).Update("online", "online-1").Update("plays", "plays+1").Error
	if err!=nil {
		log.Panicln(err)
	}
	//_, err := DB.Exec("update video_info set online = online-1 where vid = ?", vid)
	//if err != nil {
	//	fmt.Printf("update faied, error:[%v]", err.Error())
	//	return
	//}
}

//用户页视频
func UserVideo(uid string) (err error,videos []_struct.Video) {
	err = DB.Where("uid=?", uid).Find(&videos).Error
	//var sqlStr string
	//mp := make(map[int]_struct.Video)
	//
	//sqlStr = "select * from videos where uid=?"
	//row, err := DB.Query(sqlStr, uid)
	//if err != nil {
	//	return err, mp
	//}
	//defer row.Close()
	//
	//var v _struct.Video
	//
	//for i := 1; row.Next(); i++ {
	//	err = row.Scan(&v.ID, &v.UID, &v.Title, &v.Info, &v.Part, &v.VideoURL, &v.CoverURL, &v.UpTime)
	//	if err != nil {
	//		return err, mp
	//	}
	//	mp[i] = v
	//}
	return err, videos
}

func HotVideo() (err error,videos []_struct.Video) {
	err = DB.Find(&videos).Order("plays desc").Error
	//var sqlStr string
	//mp := make(map[int]_struct.Video)
	//
	//sqlStr = "select v.all from videos v,video_info i where v.id=i.vid order by i.plays DESC;"
	//row, err := DB.Query(sqlStr)
	//if err != nil {
	//	return err, mp
	//}
	//defer row.Close()
	//
	//var v _struct.Video
	////最多8个视频
	//for i := 1; row.Next() && i <= 8; i++ {
	//	err = row.Scan(&v.ID, &v.UID, &v.Title, &v.Info, &v.Part, &v.VideoURL, &v.CoverURL, &v.UpTime)
	//	if err != nil {
	//		return err, mp
	//	}
	//	mp[i] = v
	//}
	return err, videos
}
