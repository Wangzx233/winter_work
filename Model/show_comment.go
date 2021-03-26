package Model

import _struct "1/struct"

func ShowComment(vid int) (err error,mp []_struct.Comment) {
	err = DB.Where("vid=?", vid).Find(&mp).Error
	//var sqlStr string
	//mp := make(map[int]_struct.Comment)
	//
	//sqlStr ="select * from comment where vid=?"
	//row,err:=DB.Query(sqlStr,vid)
	//if err!=nil{
	//	return err,mp
	//}
	//defer row.Close()
	//
	//var cm _struct.Comment
	////
	//for i:=1;row.Next();i++ {
	//	err = row.Scan(&cm.Id,&cm.Content,&cm.Uid,&cm.Vid)
	//	if err != nil {
	//		return err,mp
	//	}
	//	mp[i]=cm
	//}
	return err,mp
}

