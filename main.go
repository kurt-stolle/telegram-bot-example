package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kurt-stolle/frank-boerman-bot/Godeps/_workspace/src/github.com/tucnak/telebot"
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

	// Create a bot
	key := os.Getenv("FRANKBOT_KEY")
	if key == "" {
		panic("Key not set, please set FRANKBOT_KEY environment variable.")
	}
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
