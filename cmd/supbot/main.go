package main

import "github.com/pxue/supbot/lib/slack"

func main() {
	token := "xoxb-19232920311-vb7KYcw8EpdfcN9Qz3v7cWpl"

	s := slack.NewClient(token)
	s.InitializeRTM()
}
