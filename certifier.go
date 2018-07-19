package main

import (
	"errors"
	"os"

	"github.com/silvnt/certifier-go/controller"
)

func main() {
	mailerKey := os.Getenv("SENDGRID_API_KEY")

	if mailerKey == "" {
		panic(errors.New("SENDGRID_API_KEY not set. Enter the value in local.env" +
			" and execute: source local.env"))
	}

	address := os.Getenv("ALTER_SERVER_ADDRESS")

	if address != "" {
		controller.StartServer(address, "home.html")
	} else {
		controller.StartServer("localhost:3000", "home.html")
	}

}
