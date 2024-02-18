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

	s := r.PathPrefix("/api").Subrouter()

	// tasks routes
	s.HandleFunc("/tasks", routes.GetTasksHandler).Methods("GET")
	s.HandleFunc("/tasks/{id}", routes.GetTaskHandler).Methods("GET")
	s.HandleFunc("/tasks", routes.CreateTaskHandler).Methods("POST")

	// users routes
	s.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	s.HandleFunc("/user/{id}", routes.GetUserHandler).Methods("GET")
	s.HandleFunc("/users", routes.CreateUsersHandler).Methods("POST")
	s.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
