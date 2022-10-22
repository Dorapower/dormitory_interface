package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:abc123321A@tcp(47.92.123.159:3306)/dormitory")
	if err != nil {
		panic(err)
	}
}
