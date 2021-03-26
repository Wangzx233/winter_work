package Model

import _struct "1/struct"

func SentBarrage(ba _struct.Barrage) error {
	err := DB.Create(_struct.Barrage{Vid: ba.Vid, Content: ba.Content}).Error
	//_,err := DB.Exec("insert barrage (vid,content) values (?,?)",ba.Vid,ba.Content)
	//if err!=nil {
	//	return err
	//}
	//
	return err
}

func ShowBarrage(vid string) (err error,barrages []_struct.Barrage) {
	err = DB.Where("vid=?", vid).Find(&barrages).Error
	//var sqlStr string
	//mp := make(map[int]_struct.Barrage)
	//
	//sqlStr ="select * from barrage where vid=?"
	//row,err:=DB.Query(sqlStr,vid)
	//if err!=nil{
	//	return err,mp
	//}
	//defer row.Close()
	//
	//var cm _struct.Barrage
	////
	//for i:=1;row.Next();i++ {
	//	err = row.Scan(&cm.Id,&cm.Content,&cm.Vid)
	//	if err != nil {
	//		return err,mp
	//	}
	//	mp[i]=cm
	//}
	return err,barrages
}