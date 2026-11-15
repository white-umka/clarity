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
	
	totalBytes := info.Totalram * uint64(info.Unit)
	freeBytes := info.Freeram * uint64(info.Unit)

	// делим на (1024 * 1024) чтобы перевести байты в мб
	usedMB := float64(totalBytes - freeBytes) / (1024 * 1024) // использовано мб памяти
	totalMb := float64(totalBytes) / (1024 * 1024) // всего мб памяти
	
	text := fmt.Sprintf("RAM: %.2f MB / %.2f MB", usedMB, totalMb)
	
	return tgbotapi.NewMessage(msg.Chat.ID, text)
}



