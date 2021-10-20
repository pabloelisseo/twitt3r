package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pabloelisseo/twitt3r/db"
	"github.com/pabloelisseo/twitt3r/models"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Error: Bad input", 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "Error: Email cannot be empty", 400)
		return
	}

	if len(u.Password) < 6 {
		http.Error(w, "Error: Password should be longer than 6 char", 400)
		return
	}

	_, userExists, _ := db.CheckUserExists(u.Email)
	if userExists {
		http.Error(w, "Error: Email already signed up", 400)
		return
	}

	_, status, err := db.InsertSignUp(u)
	if err != nil {
		http.Error(w, "Error: There was a problem inserting data from user into database"+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "Error: SignUp User not inserted into database", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
