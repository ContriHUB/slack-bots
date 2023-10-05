package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func calculate_age() {

	bot := slacker.NewClient(os.Getenv("SLACK_AGE_BOT_TOKEN"), os.Getenv("SLACK_AGE_APP_TOKEN"))

	go printCommandEvents(bot.CommandEvents())

	bot.Command("hello", &slacker.CommandDefinition{
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			response.Reply("Hi, How are you??")
		},
	})

	bot.Command("my yob is <year>", &slacker.CommandDefinition{
		Description: "yob calculator",
		Examples:    []string{"my yob is 2020"},
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				println("error")
			}
			age := 2023 - yob
			fmt.Println(age)
			r := fmt.Sprintf("age is %d", age)
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

func upload_file() {

	api := slack.New(os.Getenv("SLACK_FILE_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"files/amit.pdf", "files/capi.yaml"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}

func main() {

	os.Setenv("SLACK_FILE_BOT_TOKEN", "xoxb-4462018505568-5973904069431-w9CEyJB3Zm7CIs7bFf8GxpBQ")
	os.Setenv("SLACK_AGE_BOT_TOKEN", "xoxb-4462018505568-6007327244417-M8S1okjobscbgjQRr3GDIrR0")
	os.Setenv("SLACK_AGE_APP_TOKEN", "xapp-1-A060HKSEXRN-5980275872167-82c6e48f42620cfad148558533d7a4b7a4d8d35675a52ca59e4f34d5f034fcb6")
	os.Setenv("CHANNEL_ID", "C04D8U36JV7")
	// upload_file()
	calculate_age()
}
