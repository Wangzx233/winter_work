package Model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)
var DB *sql.DB
func MysqlInit() *sql.DB{
	db,err:=sql.Open("mysql",os.Getenv("MYSQL_DSN"))//"root:root@tcp(localhost:3306)/data"
	if err!=nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	db.SetMaxOpenConns(1000)
	err = db.Ping()
	if err != nil {
		fmt.Printf("mysql connect failed:%v", err)
		return nil
	}
	DB=db
	return DB
}

