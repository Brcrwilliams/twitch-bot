# Twitch Bot

This is a Twitch bot build in Golang using the Hellabot framework.

It was mainly created for educational purposes.

It stores commands inside a NoSQL database that is stored on the filesystem
in the project directory.

## Usage

- Create a command - `!create <command_name> <bot_reply>`
- List commands - `!list`
- Delete a command - `!remove <command_name>`

## Variables

In order to connect to Twitch, you need to have a .env file in the project directory with
the following variables set.

| Variable Name         | Description                                                                              |
| --------------------- | ---------------------------------------------------------------------------------------- |
| TWITCH_OAUTH_PASSWORD | The oauth token used to connect to Twitch chat. Get one from https://twitchapps.com/tmi/ |
| TWITCH_CHANNEL        | The name of the channel which to connect to. Ex: #selthor                                |
| TWITCH_BOT_NAME       | The username that the bot is using.                                                      |


## Start

You need to change the Mods list in main.go to include all of the users who are authorized to create / delete commands.
Then, compile a binary with:

```
go get ./...
go build .
```

You can then run the bot by executing the binary or running:

```
go run .
```