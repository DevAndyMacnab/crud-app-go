package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DevAndyMacnab/golang-course/config"
	"github.com/DevAndyMacnab/golang-course/models"
	"github.com/DevAndyMacnab/golang-course/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	port_env := os.Getenv("SERVER_PORT")
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

	fmt.Println(port_env)
	http.ListenAndServe(port_env, r)

}
