package main

import (
	"log"

	"github.com/pabloelisseo/twitt3r/db"
	"github.com/pabloelisseo/twitt3r/handlers"
)

func main() {
	err := db.CheckConnection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	handlers.Handlers()
}
