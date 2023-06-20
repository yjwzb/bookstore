package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/bookstore0612")
	if err != nil {
		fmt.Println("数据库连接失败!")
		panic(err.Error())
	}
}
