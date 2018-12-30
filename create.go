package main

import (
	"fmt"
	"regexp"
	"strings"

	hbot "github.com/whyrusleeping/hellabot"
)

// CreateCommandTrigger is used to insert a new command into the database.
var CreateCommandTrigger = hbot.Trigger{
	Condition: createCommandCondition,
	Action:    createCommandAction,
}

func createCommandCondition(b *hbot.Bot, m *hbot.Message) bool {
	if strings.HasPrefix(m.Content, "!create ") {
		for _, name := range Mods {
			if name == m.From {
				return true
			}
		}
	}
	return false
}

const createUsage = "Invalid syntax. Usage: !create <new_command_name> <new_command_response>"

// createCommandAction inserts a command into the database.
// The command should have the syntax: !create <new_command_name> <new_command_response>
func createCommandAction(b *hbot.Bot, m *hbot.Message) bool {
	args := strings.Split(m.Content, " ")
	if len(args) < 3 {
		b.Reply(m, createUsage)
		return false
	}

	name := args[1]
	response := strings.Join(args[2:], " ")

	if matched, _ := regexp.MatchString("^![a-z]+$", name); !matched {
		b.Reply(m, "Command names need to start with ! and be all lowercase letters")
		return false
	}

	command := map[string]interface{}{
		"name":     name,
		"response": response,
	}

	success := CreateCommand(command)
	if success {
		b.Reply(m, fmt.Sprintf("Created new command: %s", name))
		return true
	}
	b.Reply(m, "Failed to create command.")
	return false
}
