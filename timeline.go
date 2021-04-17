package main

import (
	"context"
	"log"

	"github.com/a-h/gemini"
	"github.com/mattn/go-mastodon"
	"github.com/mibzman/titan"
)

type TimelineViewModel struct {
	Statuses []StatusViewModel
	LastID   string
}

func createTimelineVM(statuses []*mastodon.Status) TimelineViewModel {
	var result TimelineViewModel
	for _, status := range statuses {
		result.Statuses = append(result.Statuses, createStatusVM(status))
	}
	result.LastID = string(statuses[len(statuses)-1].ID)
	return result
}

func timelineHandler(w gemini.ResponseWriter, r *gemini.Request) interface{} {
	LastID := titan.GetVar(r, "ID")
	var pg mastodon.Pagination
	if LastID != "" {
		pg.MaxID = mastodon.ID(LastID)
	}

	timeline, err := mastodonClient.GetTimelineHome(context.Background(), &pg)
	if err != nil {
		log.Fatal(err)
	}

	result := createTimelineVM(timeline)

	return result
}

func tootHandler(w gemini.ResponseWriter, r *gemini.Request) interface{} {
	ID := titan.GetVar(r, "ID")

	status, err := mastodonClient.GetStatus(context.Background(), mastodon.ID(ID))
	if err != nil {
		log.Fatal(err)
	}

	statusVM := createStatusVM(status)
	statusVM.PopulateThread()

	return statusVM
}
