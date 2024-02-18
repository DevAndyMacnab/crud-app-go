package routes

import (
	"encoding/json"
	"net/http"

	"github.com/DevAndyMacnab/golang-course/config"
	"github.com/DevAndyMacnab/golang-course/models"
	"github.com/gorilla/mux"
)

func GetUsersHandler(response http.ResponseWriter, request *http.Request) {
	var users []models.User
	config.DB.Find(&users)
	json.NewEncoder(response).Encode(&users)

}

func GetUserHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)

	var user models.User
	config.DB.First(&user, params["id"])

	if user.ID == 0 {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("User not found"))
		return
	}
	config.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	
	json.NewEncoder(response).Encode(&user)

}

func CreateUsersHandler(response http.ResponseWriter, request *http.Request) {
	var user models.User
	json.NewDecoder(request.Body).Decode(&user)
	createdUser := config.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(&user)
	//response.Write([]byte("create"))
}

func DeleteUserHandler(response http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	var user models.User

	config.DB.First(&user, params["id"])
	if user.ID == 0 {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("User not found"))
		return
	}

	config.DB.Unscoped().Delete(&user)
	response.WriteHeader(http.StatusOK)
}
