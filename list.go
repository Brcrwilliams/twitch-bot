package main

import (
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
)

// ListTrigger lists all of the commands current available.
var ListTrigger = hbot.Trigger{
	Condition: listCondition,
	Action:    listAction,
}

func listCondition(b *hbot.Bot, m *hbot.Message) bool {
	return m.Content == "!list"
}

func listAction(b *hbot.Bot, m *hbot.Message) bool {
	commands := ListCommands()
	b.Reply(m, strings.Join(commands, ", "))
	return true
}
