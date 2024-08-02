package controllers

import (
	"encoding/json"
	"myapi/models"
	"myapi/services"
	"net/http"
	"strings"
)

var userService *services.UserService

func InitUserService(service *services.UserService) {
	userService = service
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users := userService.GetAllUsers()
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Could not parse request data", http.StatusBadRequest)
			return
		}
		createdUser, err := userService.CreateUser(user)
		if err != nil {
			http.Error(w, "Could not create user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(createdUser)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	switch r.Method {
	case http.MethodGet:
		user, found := userService.GetUserByID(idStr)
		if !found {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	case http.MethodPut:
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Could not parse request data", http.StatusBadRequest)
			return
		}
		user.ID = idStr // Конвертуємо int в string
		updatedUser, err := userService.UpdateUser(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updatedUser)
	case http.MethodDelete:
		err := userService.DeleteUser(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
