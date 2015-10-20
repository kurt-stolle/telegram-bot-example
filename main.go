package main

import (
	"bytes"
	"fmt"
	"github.com/tucnak/telebot"
	"math/rand"
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
		if message.Text == "/frank" {
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

		}
	}
}
