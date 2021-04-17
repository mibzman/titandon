package main

import (
	"context"

	"github.com/a-h/gemini"
	"github.com/mattn/go-mastodon"
	"github.com/mibzman/titan"
)

func composeTootHandler(w gemini.ResponseWriter, r *gemini.Request) {
	TootContent := titan.GetQuery(r)

	toot := mastodon.Toot{}
	toot.Status = TootContent
	mastodonClient.PostStatus(context.Background(), &toot)

	w.SetHeader(gemini.CodeRedirect, "/timeline")
}

func replyTootHandler(w gemini.ResponseWriter, r *gemini.Request) {
	ID := titan.GetVar(r, "ID")
	TootContent := titan.GetQuery(r)

	toot := mastodon.Toot{}
	toot.Status = TootContent
	toot.InReplyToID = mastodon.ID(ID)
	mastodonClient.PostStatus(context.Background(), &toot)
	w.SetHeader(gemini.CodeRedirect, "/toot/"+ID)
}

func boostTootHandler(w gemini.ResponseWriter, r *gemini.Request) {
	ID := titan.GetVar(r, "ID")

	mastodonClient.Reblog(context.Background(), mastodon.ID(ID))
	w.SetHeader(gemini.CodeRedirect, "/toot/"+ID)
}

func favTootHandler(w gemini.ResponseWriter, r *gemini.Request) {
	ID := titan.GetVar(r, "ID")

	mastodonClient.Favourite(context.Background(), mastodon.ID(ID))
	w.SetHeader(gemini.CodeRedirect, "/toot/"+ID)
}
