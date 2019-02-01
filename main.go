package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"time"
)

func main() {
	loadEnv()

	bot, err := tb.NewBot(tb.Settings{
		Token:  cfg.BotToken,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	bot.Handle(tb.OnText, func(m *tb.Message) {
		response := getCatPic()

		if response.Error != nil {
			bot.Send(m.Sender, "no cat pic for u")
			return
		}

		fmt.Printf("Sending cat pic to \"%s\" (UID: %d)\n", m.Sender.FirstName, m.Sender.ID)

		pic := &tb.Photo{File: tb.FromURL(response.Url)}
		bot.Send(m.Sender, pic)
	})

	fmt.Println("Bot started")
	bot.Start()
}
