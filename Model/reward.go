package Model

import (
	_struct "1/struct"
	"log"
)

//点赞
//func Like(uid, vid string) bool {
//	err := DB.QueryRow("select * from islike where uid=? and vid=?", uid, vid).Scan() //db为sql.DB
//	if err == sql.ErrNoRows {
//		_, err := DB.Exec("insert islike (uid,vid) values (?,?)", uid, vid)
//		if err != nil {
//			fmt.Printf("Exec failed: %v", err)
//			return false
//		}
//		_, err = DB.Exec("update video_info set likes = likes+1 where vid = ?", vid)
//		if err != nil {
//			fmt.Printf("update faied, error:[%v]", err.Error())
//			return false
//		}
//		return true
//	} else {
//		return false
//	}
//}

//投币
func Coin(uid, vid, to_uid string) {
	err := DB.Model(_struct.VideoInfo{}).Where("vid=?", vid).Update("coin", "coin+1").Error
	if err != nil {
		log.Panicln(err)
		return
	}
	err = DB.Model(_struct.UserInfo{}).Where("uid=?", to_uid).Update("coin", "coin+1").Error
	if err != nil {
		log.Panicln(err)
		return
	}
	err = DB.Model(_struct.UserInfo{}).Where("uid=?", uid).Update("coin", "coin-1").Error
	if err != nil {
		log.Panicln(err)
		return
	}
	//_, err := DB.Exec("update video_info set coin = coin+1 where vid = ?",vid)
	//if err != nil {
	//	fmt.Printf("update faied, error:[%v]", err.Error())
	//	return
	//}
	//
	//
	//_, err = DB.Exec("update user_info set coin = coin+1 where uid = ?",to_uid)
	//if err != nil {
	//	fmt.Printf("update faied, error:[%v]", err.Error())
	//	return
	//}
	//
	//_, err = DB.Exec("update user_info set coin = coin-1 where uid = ?",uid)
	//if err != nil {
	//	fmt.Printf("update faied, error:[%v]", err.Error())
	//	return
	//}

}

//收藏
//func Collection(uid,vid string) bool {

//err := DB.QueryRow("select * from iscollection where uid=? and vid=?",uid,vid).Scan()   //db为sql.DB
//if err == sql.ErrNoRows {
//	//判断是否已收藏
//	_,err:=DB.Exec("insert iscollection (uid,vid) values (?,?)",uid,vid)
//	if err!=nil{
//		fmt.Printf("Exec failed: %v",err)
//		return false
//	}
//
//	//视频被收藏数加1
//	_, err = DB.Exec("update video_info set collections = collections+1 where vid = ?",vid)
//	if err != nil {
//		fmt.Printf("update faied, error:[%v]", err.Error())
//		return false
//	}
//
//	//自身收藏数加1
//	_, err = DB.Exec("update video_info set stars = stars+1 where vid = ?",vid)
//	if err != nil {
//		fmt.Printf("update faied, error:[%v]", err.Error())
//		return false
//	}
//	return true
//} else {
//	return false
//}
//}
