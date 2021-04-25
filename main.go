package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mibzman/titan"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	setupMastodon()

	server := titan.Startup()

	createPages(server)

	realCert, err := CreateCert()
	if err != nil {
		log.Fatal("error:", err)
	}

	server.Launch("localhost", realCert)

}
