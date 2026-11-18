package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"clarity/internal/config"
)

func NewBot() *tgbotapi.BotAPI {
	token := config.ReadToken()
	
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true //можно менять в ходе отладки 

	return bot
}

func StartUpdates(bot *tgbotapi.BotAPI) { 
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	for update := range bot.GetUpdatesChan(u) {
		if update.Message == nil { continue }

		if update.Message.Photo != nil {
			bot.Send(handleASCII(bot, update.Message))
			continue
		}
		
		switch update.Message.Text {
			case "/start":
				bot.Send(handleStart(bot, update.Message))
			case "/uptime":
				bot.Send(handleUptime(bot, update.Message))
			case "/freeram":
				bot.Send(handleFreeRam(bot, update.Message))
		}
	}
}

