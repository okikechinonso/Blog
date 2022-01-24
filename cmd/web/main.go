package main

import (
	"blog/Db"
	"blog/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type application struct {
	user models.IUsers
	blog models.Iblog
}

func main() {
	port := os.Getenv("PORT")
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
	db := Db.ConnectToDatabase()
	app := &application{
		user: &models.DBModel{Db: db},
		blog: &models.DBModel{Db: db},
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: app.routes(),
	}
	server.ListenAndServe()
}
