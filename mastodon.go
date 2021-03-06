package main

import (
	"context"
	"log"
	"os"

	"github.com/mattn/go-mastodon"
)

var mastodonClient *mastodon.Client

func setupMastodon() {

	c := mastodon.NewClient(&mastodon.Config{
		Server:       os.Getenv("SERVER"),
		ClientID:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
	})
	err := c.Authenticate(context.Background(), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	mastodonClient = c
}
