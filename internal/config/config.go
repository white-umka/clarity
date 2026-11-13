package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ReadToken() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("BOT_TOKEN")
	fmt.Println("токен успешно получен из .env")
	
	return token
}
