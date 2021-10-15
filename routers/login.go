package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pabloelisseo/twitt3r/bd"
	"github.com/pabloelisseo/twitt3r/jwt"
	"github.com/pabloelisseo/twitt3r/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect email or password "+err.Error(), 400)
	}
	if len(t.Email) == 0 {
		http.Error(w, "User email required "+err.Error(), 400)
	}
	user, exists := bd.TryLogin(t.Email, t.Password)
	if !exists {
		http.Error(w, "Incorrect email or password ", 400)
		return
	}
	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "There was an error while trying to generate jwt token "+err.Error(), 400)
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
