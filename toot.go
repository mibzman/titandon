package main

import (
	"context"
	"log"
	"strings"

	"github.com/LukeEmmet/html2gemini"
	"github.com/andanhm/go-prettytime"
	"github.com/mattn/go-mastodon"
)

type StatusViewModel struct {
	ID               mastodon.ID
	TimelineID       mastodon.ID
	DisplayName      string
	Account          string
	CreatedAt        string
	SpoilerText      string
	Content          string
	Status           *mastodon.Status
	Ancestors        []StatusViewModel
	Descendants      []StatusViewModel
	IsBoost          bool
	Booster          string
	MediaAttachments []MediaViewModel
}

type MediaViewModel struct {
	URL         string
	Description string
}

func tootToGmi(content string) string {
	ctx := html2gemini.NewTraverseContext(*html2gemini.NewOptions())
	text, err := html2gemini.FromString(content, *ctx)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Replace(text, "\n#", "\n #", -1)
}

func createAttachmentVMs(attachments []mastodon.Attachment) (result []MediaViewModel) {
	for _, attach := range attachments {
		result = append(result, MediaViewModel{attach.RemoteURL, attach.Description})
	}
	return
}

func createStatusVM(status *mastodon.Status) StatusViewModel {
	text := tootToGmi(status.Content)

	vm := StatusViewModel{
		status.ID,
		status.ID,
		status.Account.DisplayName,
		status.Account.Acct,
		prettytime.Format(status.CreatedAt),
		status.SpoilerText,
		text,
		status,
		[]StatusViewModel{},
		[]StatusViewModel{},
		false,
		"",
		createAttachmentVMs(status.MediaAttachments),
	}

	if status.Reblog != nil {
		vm.ID = status.Reblog.ID
		vm.DisplayName = status.Reblog.Account.DisplayName
		vm.Account = status.Reblog.Account.Acct
		vm.CreatedAt = prettytime.Format(status.Reblog.CreatedAt)
		vm.Status = status.Reblog
		vm.IsBoost = true
		vm.Booster = status.Account.DisplayName
		vm.MediaAttachments = createAttachmentVMs(status.Reblog.MediaAttachments)
	}

	return vm
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
