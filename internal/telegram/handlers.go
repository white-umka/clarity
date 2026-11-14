package telegram

import (
	"fmt"
	"syscall"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func handleStart(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "это прототип бота clarity, который был создан для отслеживания состояния сервера")
	return reply
}

func handleUptime(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	var info syscall.Sysinfo_t // переменная с структурой типа: Sysinfo_t (https://pkg.go.dev/syscall#Sysinfo_t)
	if err := syscall.Sysinfo(&info); err != nil {
		panic(err)
	}

	uptimeDuration := time.Duration(info.Uptime) * time.Second // отформатированное время

	text := fmt.Sprintf("время работы сервера: %s", uptimeDuration.String())
	
	return tgbotapi.NewMessage(msg.Chat.ID, text)
}

func handlerFreeRam(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	var info syscall.Sysinfo_t 
	if err := syscall.Sysinfo(&info); err != nil {
		panic(err)
	}

	text := fmt.Sprintf("доступная RAM: %d", info.Freeram)
	
	return tgbotapi.NewMessage(msg.Chat.ID, text)
}

