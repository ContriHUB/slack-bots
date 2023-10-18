package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
    "github.com/joho/godotenv"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"
)

func loadEnvVariables() error {
    err := godotenv.Load(".env")
    if err != nil {
        return err
    }

    return nil
}

func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvent) {
	for event := range analyticsChannel {
		fmt.Println("Command Evenets")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
	}
}

func custom_bot() {
	bot := slacker.NewClient(os.Getenv("SLACK_CUSTOM_BOT_TOKEN"), os.Getenv("SLACK_CUSTOM_APP_TOKEN"))
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
	if api == nil {
		fmt.Printf("Error creating slack client\n")
		return
	}
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"files/amit.pdf", "files/capi.yaml","files/test1.txt","files/test2.txt"}

	for i := 0; i < len(fileArr); i++ {
		  fileInfo, err := os.Stat(fileArr[i])
        if err != nil {
            fmt.Printf("File not found or error accessing file: %s\n", fileArr[i])
            continue // Skip this file and continue with the next one
        }
		      if fileInfo.IsDir() {
            fmt.Printf("Skipping directory: %s\n", fileArr[i])
            continue // Skip directories
        }
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
	loadEnvVariables()
    upload_file()
	custom_bot()
}
