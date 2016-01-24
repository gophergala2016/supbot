package main

import (
	"os"

	"github.com/gophergala2016/supbot/lib/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		panic("slack token must be set")
	}

	s := slack.NewClient(token)
	s.Start()
}
