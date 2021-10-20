package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pabloelisseo/twitt3r/db"
	"github.com/pabloelisseo/twitt3r/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid data"+err.Error(), 400)
		return
	}

	status, err := db.ModifySignUp(t, UserID)
	if err != nil {
		http.Error(w, "There was an error trying to modify the profile. Try again "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Impossible to modify user's profile. ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
