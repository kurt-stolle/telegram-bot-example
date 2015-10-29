package main

import (
	"fmt"
	"math/rand"

	"github.com/kurt-stolle/frank-boerman-bot/Godeps/_workspace/src/github.com/tucnak/telebot"
)

var frankQuotes = [8]string{"Welcome to Uni", "no lol XD", "Offcourse", "Unfortiantly", "#magicaltriangle", "indeed!", "Don't call me NSB'er!", "lol im admin, im untouchable"}

func init() {
	c := addCommand("Frank", "frank")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		fmt.Println(message.Sender.FirstName, "ran 'frank' in ", message.Chat.Title)

		if message.Sender.Username == "Flinnepin" || message.Sender.Username == "Miega" {
			bot.SendMessage(message.Chat, "🐷", nil)
			return
		}

		switch rand.Intn(6) {
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
