package main

import (
	"github.com/mibzman/titan"
)

func createPages(server titan.Server) {
	server.AddPage("/timeline", "timeline.ttn", timelineHandler)
	server.AddPage("/timeline/{ID}", "timeline.ttn", timelineHandler)
	server.AddPage("/toot/{ID}", "toot.ttn", tootHandler)

	server.AddAction("/boost/{ID}", boostTootHandler)
	server.AddAction("/fav/{ID}", favTootHandler)

	server.AddInput("/toot", "Compose your toot", composeTootHandler)
	server.AddInput("/reply/{ID}", "Reply to the toot", replyTootHandler)

}
