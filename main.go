package main

import (
	"log"

	"github.com/cemtanrikut/go-api-pills/client"
)

func main() {
	log.Println("Starting the application")
	client.MuxHandler()
}
