package main

import (
	"github.com/nlopes/slack"
	"log"
	"os"
	"github.com/paveg/redashbot/lib/match"
)

const FALSE_CODE = 1

func run(api *slack.Client) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				// log.Print("Hello Event")
				// fmt.Printf("msg: %v", msg)
			case *slack.MessageEvent:
				if match.IsTextMatch(ev.Msg.Text) {
					rtm.SendMessage(rtm.NewOutgoingMessage(ev.Msg.Text, ev.Channel))
				}

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return FALSE_CODE
			}
		}
	}
}

func main() {
	token := os.Getenv("SLACK_API_TOKEN")
	api := slack.New(token)
	os.Exit(run(api))
}
