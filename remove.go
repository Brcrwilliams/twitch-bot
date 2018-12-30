package main

import (
	"fmt"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
)

// RemoveTrigger lists all of the commands current available.
var RemoveTrigger = hbot.Trigger{
	Condition: removeCondition,
	Action:    removeAction,
}

func removeCondition(b *hbot.Bot, m *hbot.Message) bool {
	if strings.HasPrefix(m.Content, "!remove ") {
		for _, name := range Mods {
			if name == m.From {
				return true
			}
		}
	}
	return false
}

const removeUsage = "!delete must have exactly two arguments.\nUsage: !delete <command>"

func removeAction(b *hbot.Bot, m *hbot.Message) bool {
	args := strings.Split(m.Content, " ")
	if len(args) != 2 {
		b.Reply(m, removeUsage)
		return false
	}

	name := args[1]
	result := DeleteCommand(name)
	if !result {
		b.Reply(m, fmt.Sprintf("No such command: %s", name))
		return false
	}
	b.Reply(m, fmt.Sprintf("Deleted command: %s", name))
	return true
}
