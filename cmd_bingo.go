package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/tucnak/telebot"
)

type bingo struct {
	active        bool
	started       bool
	leader        string
	cards         map[string][]string
	words         []string
	gameTimer     *time.Timer
	instanceTimer *time.Timer
}

var bingos = make(map[int]*bingo)

func getBingo(id int) *bingo {

	if set, ok := bingos[id]; ok {
		set.instanceTimer.Reset(time.Hour * 12)

		return set
	}

	fmt.Println("Created a new bingo with id ", id)

	g := new(bingo)
	g.cards = make(map[string][]string)

	// Destroy if inactive for too long
	g.instanceTimer = time.NewTimer(time.Hour * 12)
	go func() {
		<-g.instanceTimer.C
		bingos[id] = nil
	}()

	bingos[id] = g

	return g
}

func init() {
	var c *command

	// Starts a new bingo
	c = addCommand("Bingo, new", "/bingo_new")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		if !message.Chat.IsGroupChat() {
			bot.SendMessage(message.Chat, "I can only do this in group chats!", nil)
			return
		}

		g := getBingo(message.Chat.ID)

		if g.active == false {
			g.active = true
			g.leader = message.Sender.Username
			g.gameTimer = time.NewTimer(time.Hour)

			go func() {
				<-g.gameTimer.C
				bot.SendMessage(message.Chat, "The bingo game has expired!", nil)
				g.active = false
				g.started = false
				g.leader = ""
				g.cards = make(map[string][]string)
				g.words = []string{}
			}()

			bot.SendMessage(message.Chat, "A new bingo was started! Commands for leader:\n/bingo_add <word> : Add a new word\n/bingo_start : Start the game\n/bingo_cancel : Stop the bingo\n/bingo_extend : Extend the bingo by 1 hour\nCommands for user:\n/bingo_card : Get a bingo card", nil)
		} else {
			bot.SendMessage(message.Chat, "There already is a bingo active", nil)
		}
	}

	// Extend the bingo
	c = addCommand("Bingo, extend", "/bingo_extend")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		if !message.Chat.IsGroupChat() {
			bot.SendMessage(message.Chat, "I can only do this in group chats!", nil)
			return
		}

		g := getBingo(message.Chat.ID)

		if g.active == true {
			if message.Sender.Username == g.leader {
				bot.SendMessage(message.Chat, "The bingo game will expire in 1 hour!", nil)
				g.gameTimer.Reset(time.Hour)
			} else {
				bot.SendMessage(message.Chat, "Nigga, you ain't the leader!", nil)
			}
		} else {
			bot.SendMessage(message.Chat, "There is no bingo active", nil)
		}
	}

	// Stop the bingo
	c = addCommand("Bingo, cancel", "/bingo_cancel")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		if !message.Chat.IsGroupChat() {
			bot.SendMessage(message.Chat, "I can only do this in group chats!", nil)
			return
		}

		g := getBingo(message.Chat.ID)

		if g.active == true {
			if message.Sender.Username == g.leader {
				bot.SendMessage(message.Chat, "The bingo game was cancelled!", nil)
				g.active = false
				g.started = false
				g.leader = ""
				g.cards = make(map[string][]string)
				g.words = []string{}
			} else {
				bot.SendMessage(message.Chat, "Nigga, you ain't the leader!", nil)
			}
		} else {
			bot.SendMessage(message.Chat, "There is no bingo active", nil)
		}
	}

	// Add a word
	c = addCommand("Bingo, add word", "/bingo_add ")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		if !message.Chat.IsGroupChat() {
			bot.SendMessage(message.Chat, "I can only do this in group chats!", nil)
			return
		}

		g := getBingo(message.Chat.ID)

		if g.active == true {
			if message.Sender.Username == g.leader {
				if g.started {
					bot.SendMessage(message.Chat, "The bingo game has already started.", nil)
					return
				}
				a := cleanCommand(message.Text, bot.Identity.Username, "/bingo_add ")

				g.words = append(g.words, a)
				bot.SendMessage(message.Chat, "Added to array!", nil)
			} else {
				bot.SendMessage(message.Chat, "Nigga, you ain't the leader!", nil)
			}
		} else {
			bot.SendMessage(message.Chat, "There is no bingo active", nil)
		}
	}

	// Start the bingo
	c = addCommand("Bingo, start", "/bingo_start")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		if !message.Chat.IsGroupChat() {
			bot.SendMessage(message.Chat, "I can only do this in group chats!", nil)
			return
		}

		g := getBingo(message.Chat.ID)

		if g.active == true {
			if message.Sender.Username == g.leader {
				if g.started {
					bot.SendMessage(message.Chat, "The bingo game has already started.", nil)
					return
				}
				bot.SendMessage(message.Chat, "The bingo game was started!\nType /bingo_card to get a card!", nil)
				g.started = true
			} else {
				bot.SendMessage(message.Chat, "Nigga, you ain't the leader!", nil)
			}
		} else {
			bot.SendMessage(message.Chat, "There is no bingo active", nil)
		}
	}

	c = addCommand("Bingo, get card", "/bingo_card")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
		if !message.Chat.IsGroupChat() {
			cardID := cleanCommand(message.Text, bot.Identity.Username, "/bingo_card ")

			flag.Parse()
			cardIDInt, err := strconv.Atoi(cardID)

			if err != nil {
				bot.SendMessage(message.Chat, "Something went wrong. I couldn't interpret your input.", nil)
				return
			}

			if g, ek := bingos[cardIDInt]; ek {
				cardID = string(message.Sender.Username)
				if card, ok := g.cards[cardID]; ok {
					fmt.Println("Giving bingo card to ", message.Sender.Username)
					bot.SendMessage(message.Chat, "Here's your card:", nil)
					for _, word := range card {
						bot.SendMessage(message.Chat, word, nil)
					}
				} else {
					fmt.Println("Card not found for ", cardID)

					bot.SendMessage(message.Chat, "I don't have that card. Did you request one in the group chat?", nil)
				}

			} else {
				bot.SendMessage(message.Chat, "I don't know that ID :(", nil)
			}

			return
		}

		id := string(message.Sender.Username)
		g := getBingo(message.Chat.ID)

		if g.started == true {
			bot.SendMessage(message.Chat, "Your ID: "+strconv.Itoa(message.Chat.ID)+"\nSend me /bingo_card <yourid> in a PM to get your card!", nil)
			fmt.Println("Card created for ", id)

			if _, ok := g.cards[id]; !ok {
				g.cards[id] = []string{}

				for i := 0; i < 5; i++ {
					g.cards[id] = append(g.cards[id], g.words[rand.Intn(len(g.words))])
				}
			}
		} else {
			bot.SendMessage(message.Chat, "There is no bingo active", nil)
		}
	}

}
