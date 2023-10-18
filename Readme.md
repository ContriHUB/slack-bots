### Prerequisites <br>
Install go v1.20+ or higher <br>
Join this slack space : [click here](https://join.slack.com/t/slack-ykk3187/shared_invite/zt-24jmn6ipy-ChgGJgobx0SCSx5mYbIp0Q)

### Setup of the project:
- Fork the repository and then clone that forked repository. <br>
- Then install all the packages by running : <br>
```go mod tidy``` <br>
- Now setup the following environment variables (Get tokens values from [here](https://drive.google.com/file/d/1EuUqV2MVh6k0OQFLUe16GQSEb3ErxRol/view?usp=sharing)):
```
SLACK_CUSTOM_APP_TOKEN=<custom app token>

SLACK_CUSTOM_BOT_TOKEN=<custom bot token>

SLACK_FILE_BOT_TOKEN=<file-bot-token>

CHANNEL_ID=<channel token>
```

- Now run the application by using :<br>
```go run main.go```
- Now go to slack space #genral channel there should be few files uploaded by the file-bot.
- You can also use custom-bot commands in #genral channel :<br>
*example* : 
```
   Input : @custom-bot my yob is 2002
   Output : your age is 21

   Input: @custom-bot hello
   Output: Hi, How are you??
```
