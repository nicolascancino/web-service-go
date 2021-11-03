package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/nicolascancino/web-service-go/server"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("error, %v", err.Error())
	}

	server.Start()

}
