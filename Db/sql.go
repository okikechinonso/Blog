package Db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "ec2-3-227-15-75.compute-1.amazonaws.com"
	port     = 5432
	user     = "toxrdqthftqgwg"
	password = "6047115930c93b6c21d63063e61de90b60ff428006ce854f83e532cb8384e758"
	dbname   = "dd7680eila7962"
)

func ConnectToDatabase() *sql.DB {
	var db *sql.DB

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected!")

	// if cfg.User == "" || cfg.DBName == ""{
	// 	return nil
	// }
	// var err error
	// db, err = sql.Open("postgres", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("conntected")
	return db
}
