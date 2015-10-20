package main

import (
	"bytes"
	"fmt"
	"github.com/tucnak/telebot"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	version = "0.0.1"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // Set the random Seed

	var frankQuotes = []string{"Welcome to Uni", "no lol XD", "Offcourse", "Unfortiantly", "#magicaltriangle"}

	fmt.Println("Starting FrankBoermanBot")

	bot, err := telebot.NewBot("170317817:AAG17sfF7EVwX65alr2XIc2CzlNfLfajqas")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bot active!")

	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		if strings.Contains(strings.ToLower(message.Text), "/frank") {
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
