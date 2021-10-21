package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/pabloelisseo/twitt3r/middlew"
	"github.com/pabloelisseo/twitt3r/routers"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlew.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/showprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ShowProfile))).Methods("GET")
	router.HandleFunc("/modifyprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/readtweets", middlew.CheckDB(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/deletetweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadavatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getavatar", middlew.CheckDB(routers.GetAvatar)).Methods("GET")
	router.HandleFunc("/uploadbanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getbanner", middlew.CheckDB(routers.GetBanner)).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
