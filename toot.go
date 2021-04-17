package main

import (
	"context"
	"log"

	"github.com/LukeEmmet/html2gemini"
	"github.com/andanhm/go-prettytime"
	"github.com/mattn/go-mastodon"
)

type StatusViewModel struct {
	ID          mastodon.ID
	DisplayName string
	Account     string
	CreatedAt   string
	SpoilerText string
	Content     string
	Status      *mastodon.Status
	Ancestors   []StatusViewModel
	Descendants []StatusViewModel
}

func createStatusVM(status *mastodon.Status) StatusViewModel {
	ctx := html2gemini.NewTraverseContext(*html2gemini.NewOptions())
	text, err := html2gemini.FromString(status.Content, *ctx)
	if err != nil {
		log.Fatal(err)
	}

	return StatusViewModel{
		status.ID,
		status.Account.DisplayName,
		status.Account.Acct,
		prettytime.Format(status.CreatedAt),
		status.SpoilerText,
		text,
		status,
		[]StatusViewModel{},
		[]StatusViewModel{},
	}
}

func (status *StatusViewModel) PopulateThread() {
	thread, err := mastodonClient.GetStatusContext(context.Background(), status.ID)
	if err != nil {
		log.Fatal(err)
	}

	for _, Ancestor := range thread.Ancestors {
		status.Ancestors = append(status.Ancestors, createStatusVM(Ancestor))
	}

	for _, Decendant := range thread.Descendants {
		status.Descendants = append(status.Descendants, createStatusVM(Decendant))
	}
}
