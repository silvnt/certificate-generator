package main

import (
	"github.com/silvnt/certifier-go/controller"
)

func main() {
	controller.StartServer("localhost:3000", "home.html")
}
