package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Comment Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-386283242208-4427916161011-xiipO1q9vgyVC2Nkw2gpDJuM")
	os.Setenv("SLACK_APP_TOKEN", "xapp-1-A04CL05FMJ6-4425060832821-6fe942c32e846da44c93f083e392d4d98f01672d98dc947c3047c522bdfc8ebc")


	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("My YOB is <year>", &slacker.CommandDefinition{
		Description: "YOB Calculator",
		// Example: "My YOB is 2020",
		Handler: func (botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter)  {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)

			if err != nil {
				println("error")
			}

			age := 2022-yob
			r := fmt.Sprintf("Your age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}