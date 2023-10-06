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
	}
}

func calculate_age() {
	bot := slacker.NewClient(os.Getenv("SLACK_AGE_BOT_TOKEN"), os.Getenv("SLACK_AGE_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())
	bot.Command("Hello", &slacker.CommandDefinition{
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
	os.Setenv("SLACK_FILE_BOT_TOKEN", "xoxb-4462018505568-5973904069431-EYD1DuOVBDobrCajM2PysBRA")
	os.Setenv("SLACK_AGE_BOT_TOKEN", "xoxb-4462018505568-6007327244417-gbVaG24z59Ov57pW9cU7MZYk")
	os.Setenv("SLACK_AGE_APP_TOKEN", "xapp-1-A060HKSEXRN-5992133251221-dee034e996d29fe2d23dc5a7d944b58a043d8f651754c63c9ab9c3d14386f8e8")
	os.Setenv("CHANNEL_ID", "C04D8U36JV7")
	upload_file()
	calculate_age()
}
