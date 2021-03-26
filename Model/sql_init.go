package Model

import (
	_struct "1/struct"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

//func MysqlInit() *sql.DB{
//	db,err:=sql.Open("mysql",os.Getenv("MYSQL_DSN"))//"root:root@tcp(localhost:3306)/data"
//	if err!=nil {
//		fmt.Printf("mysql connect failed:%v", err)
//		return nil
//	}
//	db.SetMaxOpenConns(1000)
//	err = db.Ping()
//	if err != nil {
//		fmt.Printf("mysql connect failed:%v", err)
//		return nil
//	}
//	DB=db
//	return DB
//}
var DB *gorm.DB

const (
	dns = "root:root@tcp(127.0.0.1:3306)/data1?charset=utf8mb4&parseTime=True&loc=Local"
)
func SqlInit() {
	db, err := gorm.Open(mysql.Open(dns),&gorm.Config{})
	//db = db.Debug()
	if err != nil {
		log.Panicln(err)
		return
	}
	if db.AutoMigrate(&_struct.Comment{},&_struct.UserInfo{},&_struct.Video{},&_struct.VideoInfo{},&_struct.User{},&_struct.Barrage{}) != nil {
		log.Println(err)
		return
	}
	DB = db
}

