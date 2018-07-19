package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/silvnt/certifier-go/controller"
)

func main() {
	mailerKey := os.Getenv("SENDGRID_API_KEY")

	if mailerKey == "" {
		panic(errors.New("SENDGRID_API_KEY not set. Enter the value in local.env" +
			" and execute: source local.env"))
	}

	address := os.Getenv("SERVER_ADDRESS")

	controller.HandleHome("templates/home.html")
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))

	if address != "" {
		log.Println("Listening " + address)
		http.ListenAndServe(address, nil)

	} else {
		log.Println("Listening localhost:3000")
		http.ListenAndServe("localhost:3000", nil)
	}

}
