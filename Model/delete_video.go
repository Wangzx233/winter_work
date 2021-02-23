package Model

import "fmt"

func DeleteVideo(uid,vid string) error {

	stmt, err := DB.Prepare("DELETE FROM videos WHERE id=? and uid=?")
	if err!=nil{
		return err
	}
	defer stmt.Close().Error()

	stmt1,err := DB.Prepare("DELETE FROM video_info WHERE vid=? and uid=?")
	if err!=nil{
		return err
	}

	res, err := stmt.Exec(vid,uid)
	if err!=nil{
		return err
	}
	num, err := res.RowsAffected()
	if err!=nil{
		return err
	}


	res1, err := stmt1.Exec(vid,uid)
	if err!=nil{
		return err
	}
	num1, err := res1.RowsAffected()
	if err!=nil{
		return err
	}

	fmt.Println(num,num1)





	return err
}
