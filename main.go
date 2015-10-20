package main

import (
	"bytes"
	"fmt"
	"github.com/tucnak/telebot"
	"io"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	version = "0.0.1"
)

func main() {
	// Hello world
	fmt.Println("Starting FrankBoermanBot...")

	// Basic variables
	rand.Seed(time.Now().UTC().UnixNano()) // Set the random Seed

	var frankQuotes = []string{"Welcome to Uni", "no lol XD", "Offcourse", "Unfortiantly", "#magicaltriangle"}

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
			panic(err)
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := bufKey.Write(bufRead[:n]); err != nil {
			panic(err)
		}
	}

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
		if strings.Contains(strings.ToLower(message.Text), "/frank") {
			fmt.Println(message.Sender.FirstName, "ran the /frank command...")
			switch rand.Intn(5) {
			case 1:
				var buffer bytes.Buffer

				buffer.WriteString("STFU ")
				buffer.WriteString(message.Sender.FirstName)
				buffer.WriteString("!")

				bot.SendMessage(message.Chat, buffer.String(), nil)
				break
			case 2:
				var buffer bytes.Buffer

				buffer.WriteString("NO SPAM! Last warning ")
				buffer.WriteString(message.Sender.FirstName)
				buffer.WriteString("!")

				bot.SendMessage(message.Chat, buffer.String(), nil)
				break
			default:
				bot.SendMessage(message.Chat,
					frankQuotes[rand.Intn(len(frankQuotes))], nil)
			}

		} else if strings.Contains(strings.ToLower(message.Text), "/magictriangle") {
			fmt.Println(message.Sender.FirstName, "ran the /magictriangle command...")
			query := strings.Replace(strings.ToLower(message.Text), "/magictriangle", "", -1)
			query = url.QueryEscape(query)

			bot.SendMessage(message.Chat, "Use the #magictriangle!", nil)

			var bufGoogle bytes.Buffer
			bufGoogle.WriteString("https://www.google.com/search?q=")
			bufGoogle.WriteString(query)
			bot.SendMessage(message.Chat, bufGoogle.String(), nil)

			var bufSO bytes.Buffer
			bufSO.WriteString("http://stackoverflow.com/search?q=")
			bufSO.WriteString(query)
			bot.SendMessage(message.Chat, bufSO.String(), nil)

			var bufWiki bytes.Buffer
			bufWiki.WriteString("https://en.wikipedia.org/w/index.php?search=")
			bufWiki.WriteString(query)
			bot.SendMessage(message.Chat, bufWiki.String(), nil)
		}
	}
}
