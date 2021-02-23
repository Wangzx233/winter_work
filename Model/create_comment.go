package Model

import (
	_struct "1/struct"
	"fmt"
)

func CreateComment(com _struct.Comment) error {

	_,err:=DB.Exec("insert comment (content,uid,vid) values (?,?,?)",com.Content,com.Uid,com.Vid)
	if err!=nil{
		fmt.Printf("Exec failed: %v",err)
		return err
	}
	return err
}
