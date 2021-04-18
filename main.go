package main

import (
	"log"

	"github.com/mibzman/titan"
)

func main() {

	setupMastodon()

	server := titan.Startup()

	createPages(server)

	realCert, err := CreateCert()
	if err != nil {
		log.Fatal("error:", err)
	}

	server.Launch("localhost", realCert)

}
