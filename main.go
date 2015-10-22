package main

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/tucnak/telebot"
)

const (
	version = "0.0.1"
)

type command struct {
	name      string
	initiator string
	action    func(bot *telebot.Bot, message *telebot.Message)
}

var commands []*command

func addCommand(name string, initiator string) *command {
	var c = new(command)
	c.name = name
	c.initiator = initiator

	commands = append(commands, c)

	fmt.Println("Loaded command: ", c.initiator)

	return c
}

func cleanCommand(s string, u string, c string) string {
	var query string
	query = strings.ToLower(strings.Replace(s, "@"+u, "", -1))
	query = strings.Replace(query, c, "", -1)

	return query
}

func main() {
	// Hello world
	fmt.Println("Starting Frank Boerman Bot...")

	// Basic variables
	rand.Seed(time.Now().UTC().UnixNano()) // Set the random Seed

	// Read the key
	var bufKey bytes.Buffer
	fi, err := os.Open("key.txt") // Key is in current directory
	if err != nil {               // Check for errors
		panic(err)
	}

	defer func() { // close fi on exit and check for its returned error
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	bufRead := make([]byte, 1024) // Make a buffer to read into
	for {
		// read a chunk
		n, err := fi.Read(bufRead)
		if err != nil && err != io.EOF {
			fi.Close()
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := bufKey.Write(bufRead[:n]); err != nil {
			fi.Close()
			panic(err)
		}
	}

	fi.Close()

	// Create a bot
	key := bufKey.String()
	fmt.Println("Telegram API key:", key)
	bot, err := telebot.NewBot(key)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Start listenting
	fmt.Println(bot.Identity.FirstName, " is active")

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		text := strings.ToLower(strings.Replace(message.Text, "@"+bot.Identity.Username, "", -1))

		for _, c := range commands {
			if strings.Contains(text, c.initiator) {
				c.action(bot, &message)
			}
		}
	}
}
