package main

import (
	"log"

	"github.com/mibzman/titan"
)

func main() {

	setupMastodon()

	server := titan.GenerateServer()

	createPages(server)

	realCert, err := CreateCert()
	if err != nil {
		log.Fatal("error:", err)
	}

	server.Launch("localhost", realCert)

}
