package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/pabloelisseo/twitt3r/db"
)

func GetBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter id", http.StatusBadRequest)
		return
	}

	profile, err := db.FindProfile(ID)
	if err != nil {
		http.Error(w, "You must send the parameter id", http.StatusBadRequest)
		return
	}

	file, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Cannot find image", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Cannot copy image", http.StatusBadRequest)
		return
	}
}
