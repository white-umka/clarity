package config

import (
	"log"
	"os"
)

func ReadToken() string {
	token := os.Getenv("BOT_TOKEN")
	if token == "my_token" {
		log.Fatal("ОШИБКА: не задана переменная окружения BOT_TOKEN")
	}
	return token
}