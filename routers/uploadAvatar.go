package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/pabloelisseo/twitt3r/db"
	"github.com/pabloelisseo/twitt3r/models"
)

func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	if err != nil {
		http.Error(w, "You must send the avatar field! "+err.Error(), http.StatusBadRequest)
		return
	}
	var extension = strings.Split(handler.Filename, ".")[1]
	var filePath string = "uploads/avatars/" + UserID + "." + extension

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error uploading the image! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error copying the image! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = UserID + "." + extension
	status, err = db.ModifySignUp(user, UserID)
	if err != nil || !status {
		http.Error(w, "Error storing the image into database! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
