package routers

import (
	"net/http"

	"github.com/pabloelisseo/twitt3r/db"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter id", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, UserID)
	if err != nil {
		http.Error(w, "There was an error trying to delete the tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
