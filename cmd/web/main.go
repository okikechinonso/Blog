package main

import (
	"blog/Db"
	"blog/models"
	"net/http"
)

type application struct {
	user models.IUsers
	blog models.Iblog
}

func main() {
	db := Db.ConnectToDatabase()
	app := &application{
		user: &models.DBModel{Db: db},
		blog: &models.DBModel{Db: db},
	}

	server := &http.Server{
		Addr:    "localhost:5000",
		Handler: app.routes(),
	}
	server.ListenAndServe()
}
