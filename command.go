package main

import (
	hbot "github.com/whyrusleeping/hellabot"
)

// CommandTrigger reads a command from the database and replies with the response.
var CommandTrigger = hbot.Trigger{
	Condition: commandCondition,
	Action:    commandAction,
}

func commandCondition(b *hbot.Bot, m *hbot.Message) bool {
	return CommandExists(m.Content)
}

func commandAction(b *hbot.Bot, m *hbot.Message) bool {
	cmd := GetCommand(m.Content)
	b.Reply(m, cmd.Response)
	return true
}
