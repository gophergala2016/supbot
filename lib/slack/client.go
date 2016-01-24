package slack

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/gophergala2016/supbot/Godeps/_workspace/src/github.com/nlopes/slack"
	"github.com/gophergala2016/supbot/lib/hal"
)

// Slack represets a Slack bot.
type Slack struct {
	token  string // slack token
	rtm    *slack.RTM
	api    *slack.Client
	botUID string

	// singleton channel name
	channel string
}

var (
	supBot io.Writer
)

// NewClient creates a new slack bot.
func NewClient(token string) *Slack {
	if len(token) < 1 {
		panic("supbot: can't seem to start myself")
	}
	api := slack.New(token)

	s := &Slack{token: token, api: api, rtm: api.NewRTM()}
	go s.rtm.ManageConnection()

	supBot = hal.New(s)
	return s
}

func (s *Slack) wasMentioned(msg string) bool {
	if len(msg) < 1 {
		return false
	}
	// NOTE: must be prefixed
	return strings.HasPrefix(msg, s.botUID)
}

// expect some byte and write to slack
func (s *Slack) Write(o []byte) (n int, err error) {
	outBuf := bytes.Buffer{}
	outBuf.Write(o)

	s.rtm.SendMessage(
		s.rtm.NewOutgoingMessage(
			outBuf.String(),
			s.channel,
		),
	)
	//params := slack.NewPostMessageParameters()
	//params.Username = "supslack"
	//params.AsUser = true

	//params.Attachments = []slack.Attachment{
	//{
	////AuthorName: authorName,
	////AuthorIcon: authorIcon,
	//Text: fmt.Sprintf("%s\n\u2014\n", outBuf.String()), // \u200B for space
	////ThumbURL: thumbURL,
	////Fields: []AttachmentField{
	////AttachmentField{
	////Title: "",
	////Value: fmt.Sprintf("*<%s|%s>*\n%s", post.ShortUrl, post.Title, description),
	////Short: false,
	////},
	////},
	//MarkdownIn: []string{"text"},
	//}}

	//s.api.PostMessage(s.channel, "", params)
	return len(o), nil
}

// Start waits for Slack events.
func (s *Slack) Start() {
Loop:
	for {
		select {
		case msg := <-s.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Println("slackbot: hello dave.")
			case *slack.ConnectedEvent:
				log.Println("slackbot: I'm online dave.")
				for _, ch := range ev.Info.Channels {
					log.Printf("slackbot: joined channel %s\n", ch.Name)
					s.rtm.SendMessage(
						s.rtm.NewOutgoingMessage(
							ch.Name,
							"Never send a human to do a machine's job.",
						),
					)
				}
				s.botUID = fmt.Sprintf("<@%s>: ", ev.Info.User.ID)
			case *slack.MessageEvent:
				s.channel = ev.Msg.Channel
				// must be mentioned
				if s.wasMentioned(ev.Text) {
					// TODO: pass to bot
					supBot.Write([]byte(strings.TrimPrefix(ev.Text, s.botUID)))
				}
			case *slack.InvalidAuthEvent:
				log.Println("supbot: I seem to be disconnected, can't let you do that.")
				break Loop
			}
		}
	}
}
