package Db

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)



func ConnectToDatabase() *sql.DB{
	var db *sql.DB
	cfg := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "CMS",
	}
	if cfg.User == "" || cfg.DBName == ""{
		return nil
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("conntected")
	return db
}
