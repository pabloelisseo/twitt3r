package main

import (
	"log"

	"github.com/pabloelisseo/twitt3r/bd"
	"github.com/pabloelisseo/twitt3r/handlers"
)

func main() {
	err := bd.CheckConnection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	handlers.Handlers()
}
