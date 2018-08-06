package main

import (
	"os"
	"log"

	"github.com/nlopes/slack"
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
				log.Print("Hello Event")

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

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
