package Model

import _struct "1/struct"

func SentBarrage(ba _struct.Barrage) error {
	_,err := DB.Exec("insert barrage (uid,vid,content) values (?,?,?)",ba.Uid,ba.Vid,ba.Content)
	if err!=nil {
		return err
	}

	return err
}

func ShowBarrage(vid string) (error,map[int]_struct.Barrage) {
	var sqlStr string
	mp := make(map[int]_struct.Barrage)

	sqlStr ="select * from barrage where vid=?"
	row,err:=DB.Query(sqlStr,vid)
	if err!=nil{
		return err,mp
	}
	defer row.Close()

	var cm _struct.Barrage
	//
	for i:=1;row.Next();i++ {
		err = row.Scan(&cm.Id,&cm.Content,&cm.Uid,&cm.Vid)
		if err != nil {
			return err,mp
		}
		mp[i]=cm
	}
	return err,mp
}