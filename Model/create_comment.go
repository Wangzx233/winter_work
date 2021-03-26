package Model

import (
	_struct "1/struct"

)

func CreateComment(com _struct.Comment) error {
	err := DB.Create(&_struct.Comment{Content: com.Content, Uid: com.Uid, Vid: com.Vid}).Error
	//_,err:=DB.Exec("insert comment (content,uid,vid) values (?,?,?)",com.Content,com.Uid,com.Vid)
	//if err!=nil{
	//	fmt.Printf("Exec failed: %v",err)
	//	return err
	//}
	return err
}
