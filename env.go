package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type config struct {
	BotToken string
}

var cfg config

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		os.Exit(0)
	}

	cfg.BotToken = os.Getenv("BOT_TOKEN")

	if cfg.BotToken == "" {
		log.Println("No BOT_TOKEN in .env file")
		os.Exit(0)
	}
}
