package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pabloelisseo/twitt3r/db"
	"github.com/pabloelisseo/twitt3r/models"
)

func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Invalid tweet.", 400)
		return
	}

	tweet := models.RecordTweet{
		UserId:    UserID,
		Message:   message.Message,
		CreatedAt: time.Now(),
	}

	_, status, err := db.InsertTweet(tweet)
	if err != nil {
		http.Error(w, "There was an error trying to insert the tweet. Try again"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "Impossible to insert the tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
