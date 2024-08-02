package controllers

import (
	"encoding/json"
	"myapi/models"
	"myapi/pkg/authentication_check"
	"myapi/pkg/token"
	"myapi/storage"
	"net/http"
)

type AuthHandler struct {
    UserStorage *storage.UserStorage
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var usr models.User
    err := json.NewDecoder(r.Body).Decode(&usr)
    if err != nil {
        http.Error(w, "Could not parse request data", http.StatusBadRequest)
        return
    }

    err = authentication_check.ValidateCredentials(&usr, h.UserStorage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    token, err := token.GenerateToken(usr.Email, usr.Role, usr.ID)
    if err != nil {
        http.Error(w, "Problem with generating a token", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{
        "message": "Login successful",
        "token":   token,
    })
}
