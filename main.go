package main

import (
	"fmt"
	"github.com/tucnak/telebot"
	"time"
)

const (
	version = "0.0.1"
)

func main() {
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
			bot.SendMessage(message.Chat,
				"Hello, "+message.Sender.FirstName+"!", nil)
		}
	}
}
