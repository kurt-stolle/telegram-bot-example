package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/kurt-stolle/frank-boerman-bot/Godeps/_workspace/src/github.com/tucnak/telebot"
)

func init() {
	c := addCommand("Magic Triangle", "/magictriangle")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		fmt.Println(message.Sender.FirstName, "ran /magictriangle in ", message.Chat.Title)

		var query string
		query = strings.ToLower(strings.Replace(message.Text, "@"+bot.Identity.Username, "", -1))
		query = strings.Replace(query, "/magictriangle", "", -1)
		query = url.QueryEscape(query)

		bot.SendMessage(message.Chat, "Use the #magictriangle!\nhttps://www.google.com/search?q="+query+"\nhttp://stackoverflow.com/search?q="+query+"\nhttps://en.wikipedia.org/w/index.php?search="+query, nil)
	}
}
