package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// InitDB 初始化数据库
func InitDB() (err error) {
	Db, err = sqlx.Connect("mysql", "root:root@tcp(192.168.44.10:3306)/golang?parseTime=true")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(5)
	return
}
