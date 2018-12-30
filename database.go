package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/HouzuoGuo/tiedot/db"
	log "gopkg.in/inconshreveable/log15.v2"
)

var conn *db.Col

// Command represents a Twitch chat command.
type Command struct {
	Name     string `json:"name"`
	Response string `json:"response"`
}

func initDatabase() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	databaseDirectory := filepath.Join(cwd, "database")
	log.Info(fmt.Sprintf("Opening database at: %s", databaseDirectory))
	commandDatabase, err := db.OpenDB(databaseDirectory)
	if err != nil {
		return err
	}

	collection := "Commands"
	if !collectionExists(commandDatabase, collection) {
		log.Info(fmt.Sprintf("Creating Collection: %s", collection))
		if err := commandDatabase.Create(collection); err != nil {
			return err
		}
	}

	conn = commandDatabase.Use(collection)
	return nil
}

func collectionExists(db *db.DB, collectionName string) bool {
	for _, name := range db.AllCols() {
		if name == collectionName {
			return true
		}
	}
	return false
}

// CreateCommand inserts a command into the NoSQL database.
func CreateCommand(cmd map[string]interface{}) bool {
	_, err := conn.Insert(cmd)
	if err != nil {
		log.Error(fmt.Sprintf("Could not insert command: %s", cmd["name"]))
		return false
	}
	return true
}

// CommandExists checks if the command with the given name
// exists in the database.
func CommandExists(name string) (result bool) {
	result = false
	conn.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		cmd := new(Command)
		json.Unmarshal(doc, cmd)
		if cmd.Name == name {
			result = true
			return false
		}
		return true
	})
	return result
}

// GetCommand reads the command with the given name from the database.
func GetCommand(name string) (result *Command) {
	conn.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		cmd := new(Command)
		json.Unmarshal(doc, cmd)
		if cmd.Name == name {
			result = cmd
			return false
		}
		return true
	})
	return result
}

// ListCommands returns a slice containing the names of all commands
// that exist in the database.
func ListCommands() (list []string) {
	conn.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		cmd := new(Command)
		json.Unmarshal(doc, cmd)
		list = append(list, cmd.Name)
		return true
	})
	return list
}

// getCommandIndex finds the index of a command with the given name.
func getCommandIndex(name string) (index int) {
	success := false
	conn.ForEachDoc(func(id int, doc []byte) (moveOn bool) {
		cmd := new(Command)
		json.Unmarshal(doc, cmd)
		if cmd.Name == name {
			index = id
			success = true
			return false
		}
		return true
	})
	if !success {
		return -1
	}
	return index
}

// DeleteCommand deletes the command with the given name.
// It returns false if the command does not exist.
func DeleteCommand(name string) bool {
	index := getCommandIndex(name)
	if index == -1 {
		return false
	}
	conn.Delete(index)
	return true
}
