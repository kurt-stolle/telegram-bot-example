package main

import (
	"bytes"
	"fmt"
	"github.com/tucnak/telebot"
	"io"
	"math/rand"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	version = "0.0.1"
)

func main() {
	// Hello world
	fmt.Println("Starting Frank Boerman Bot...")

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
		text := strings.Replace(strings.ToLower(message.Text), "@"+bot.Identity.Username, "", -1)

		if strings.Contains(text, "frank") {
			fmt.Println(message.Sender.FirstName, "ran 'frank' in ", message.Chat.Title)
			switch rand.Intn(5) {
			case 1:
				bot.SendMessage(message.Chat, "STFU "+message.Sender.FirstName+"!", nil)
				break
			case 2:
				bot.SendMessage(message.Chat, "NO SPAM! Last warning "+message.Sender.FirstName+"!", nil)
				break
			default:
				bot.SendMessage(message.Chat,
					frankQuotes[rand.Intn(len(frankQuotes))], nil)
			}

		} else if strings.Contains(text, "/magictriangle") {
			fmt.Println(message.Sender.FirstName, "ran /magictriangle in ", message.Chat.Title)

			var query string
			query = strings.Replace(text, "/magictriangle", "", -1)
			query = url.QueryEscape(query)

			bot.SendMessage(message.Chat, "Use the #magictriangle!\nhttps://www.google.com/search?q="+query+"\nhttp://stackoverflow.com/search?q="+query+"\nhttps://en.wikipedia.org/w/index.php?search="+query, nil)
		} else if strings.Contains(text, "/essay") {
			fmt.Println(message.Sender.FirstName, "ran /essay in ", message.Chat.Title)

			dirname := "." + string(filepath.Separator) + "essays" + string(filepath.Separator)
			d, err := os.Open(dirname)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer d.Close()

			files, err := d.Readdir(-1)

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			var validFiles []os.FileInfo

			for _, file := range files {
				if file.Mode().IsRegular() {
					validFiles = append(validFiles, file)
				}
			}

			var randomEssay = validFiles[rand.Intn(len(validFiles))]

			var buf bytes.Buffer
			fi, err := os.Open("essays" + string(filepath.Separator) + randomEssay.Name()) // Key is in current directory
			if err != nil {                                                                // Check for errors
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
				if _, err := buf.Write(bufRead[:n]); err != nil {
					fi.Close()
					panic(err)
				}
			}

			fi.Close()

			bot.SendMessage(message.Chat, buf.String(), nil)
		}
	}
}
