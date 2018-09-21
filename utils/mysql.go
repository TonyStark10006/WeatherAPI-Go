package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		log.Fatal(err)
	}
	return
}

func QueryAll() {

}
