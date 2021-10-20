package routers

import (
	"encoding/json"
	"net/http"

	"github.com/pabloelisseo/twitt3r/db"
)

func ShowProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "ID parameter missing", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)
	if err != nil {
		http.Error(w, "Cannot find user profile"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)

}
