package main

import (
	"io"

	"github.com/pxue/supbot/slack"
)

var (
	Supbot io.Writer
)

func main() {
	token := "xoxb-19232920311-vb7KYcw8EpdfcN9Qz3v7cWpl"

	s := slack.NewClient(token)
	s.InitializeRTM()

	Supbot = NewHal(s)
}
