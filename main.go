package main

import (
	"net/http"

	"github.com/DevAndyMacnab/golang-course/config"
	"github.com/DevAndyMacnab/golang-course/models"
	"github.com/DevAndyMacnab/golang-course/routes"
	"github.com/gorilla/mux"
)

func main() {

	config.DataBaseConnection()
	config.DB.AutoMigrate(models.Task{})
	config.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	http.ListenAndServe(":3000", r)
}
