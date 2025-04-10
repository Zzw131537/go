package main

import (
	"database/sql"
	"log"
	// _ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "gee.db")
	defer func() {
		_ = db.Close()
	}()
	if err == nil {
		db.Exec("drop table if exists User;")
		db.Exec("create table user(Name text);")
		result, err := db.Exec("insert into User(`Name`) values (?),(?)", "Tom", "Sam")

		if err == nil {
			affected, _ := result.RowsAffected()
			log.Println(affected)
		}
		row := db.QueryRow("select Name from User limit 1")
		var name string
		if err := row.Scan(&name); err == nil {
			log.Println(name)
		}
	} else {
		log.Println("数据库连接失败!")
	}
}
