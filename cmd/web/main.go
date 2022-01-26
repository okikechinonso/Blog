package main

import (
	"blog/Db"
	"blog/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	port := os.Getenv("PORT")
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
	}
	db := Db.ConnectToDatabase()
	app := &Application{
		user: &models.DBModel{Db: db},
		blog: &models.DBModel{Db: db},
	}
	r := app.routes()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
	server.ListenAndServe()
}
