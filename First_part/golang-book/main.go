package main

import (
    "log"
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
)
// Запуск датабазы sql

func main() {
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
    err = db.Ping()
if err != nil {
	// do something here
}//проверка на работоспособность
	defer db.Close()
}