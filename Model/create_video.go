package Model

import (
	_struct "1/struct"
)

func CreateVideo(video _struct.Video) error {
	err := DB.Create(&_struct.Video{
		UID:         video.UID,
		Title:       video.Title,
		Info:        video.Info,
		Part:        video.Part,
		VideoURL:    video.VideoURL,
		CoverURL:    video.CoverURL,
		VideoInfo:   _struct.VideoInfo{UID: video.UID},
		VideoInfoID: 0,
	}).Error
	//res,err := DB.Exec("insert  videos (uid,title,info,part,videourl,coverurl) values (?,?,?,?,?,?)", video.UID,video.Title,video.Info,video.Part,video.VideoURL,video.CoverURL)
	//if err!=nil {
	//	return err
	//}
	//id,err := res.LastInsertId()
	//if err!=nil {
	//	return err
	//}
	//_,err = DB.Exec("insert  video_info (vid,uid) values (?,?)",id,video.UID)
	//if err!=nil {
	//	return err
	//}
	return err
}

