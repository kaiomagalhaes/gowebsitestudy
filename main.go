package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gowebsitestudy/controllers"
	"gowebsitestudy/models"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "usegolang_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	ug, err := models.NewUserGorm(psqlInfo)
	if err != nil {
		panic(err)
	}

	ug.AutoMigrate()

	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(ug)

	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}
