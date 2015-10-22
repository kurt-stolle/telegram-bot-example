package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/tucnak/telebot"
)

func init() {
	c := addCommand("Essay", "/essay")
	c.action = func(bot *telebot.Bot, message *telebot.Message) {
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

		d.Close()

		s, err := ioutil.ReadFile("essays" + string(filepath.Separator) + randomEssay.Name()) // Key is in current directory
		if err != nil {                                                                       // Check for errors
			panic(err)
		}

		bot.SendMessage(message.Chat, string(s), nil)
	}
}
