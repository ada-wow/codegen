package main

import (
	"codegen/mysqldb/Build/DAL/DBUtility"
	"database/sql"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	db, _ := sql.Open("mysql", sqlhelper.GetConnectionString())
	stmt, _ := db.Prepare("INSERT INTO person(username,age,address) VALUES(?,?,?)")
	_, err := stmt.Exec("ada", 2, "homes")
	if err != nil {
		panic(err)
	}
}
