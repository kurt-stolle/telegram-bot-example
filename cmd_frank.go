package main

import (
	"fmt"
	"math/rand"

	"github.com/tucnak/telebot"
)

var frankQuotes = [5]string{"Welcome to Uni", "no lol XD", "Offcourse", "Unfortiantly", "#magicaltriangle"}

func init() {
	c := addCommand("Frank", "frank")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		fmt.Println(message.Sender.FirstName, "ran 'frank' in ", message.Chat.Title)
		switch rand.Intn(5) {
		case 1:
			bot.SendMessage(message.Chat, "STFU "+message.Sender.FirstName+"!", nil)
			break
		case 2:
			bot.SendMessage(message.Chat, "NO SPAM! Last warning "+message.Sender.FirstName+"!", nil)
			break
		default:
			bot.SendMessage(message.Chat, frankQuotes[rand.Intn(len(frankQuotes))], nil)
		}
	}
}
