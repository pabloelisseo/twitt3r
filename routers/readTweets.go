package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pabloelisseo/twitt3r/db"
)

func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the parameter page", http.StatusBadRequest)
		return
	}

	pageString, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the parameter page with a value greater than 0", http.StatusBadRequest)
		return
	}

	page := int64(pageString)
	response, correct := db.ReadTweet(ID, page)
	if correct == false {
		http.Error(w, "Error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
