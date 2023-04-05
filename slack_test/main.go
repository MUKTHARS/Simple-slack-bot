package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()

	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5059182069684-5050228543686-slH00gb9O53sx1tHUQUzRsjr")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A05186TQ5CP-5056641983715-8fc4169741e57f83fbf90e3aa6f44f9382c6d563a4d7905c44bb7e4d8e79c573")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT"), os.Getenv("SLACK_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("pong")
		},
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
