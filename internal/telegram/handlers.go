package telegram

import (
	"fmt"
	"syscall"
	"time"
	"os"
	"bufio"
	"strings"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func handleStart(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	reply := tgbotapi.NewMessage(msg.Chat.ID, "это прототип бота clarity, который был создан для отслеживания состояния сервера")
	return reply
}

func handleUptime(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	var info syscall.Sysinfo_t // переменная с стру00ктурой типа: Sysinfo_t (https://pkg.go.dev/syscall#Sysinfo_t)
	if err := syscall.Sysinfo(&info); err != nil {
		panic(err)
	}

	uptimeDuration := time.Duration(info.Uptime) * time.Second // отформатированное время

	text := fmt.Sprintf("время работы сервера: %s", uptimeDuration.String())
	
	return tgbotapi.NewMessage(msg.Chat.ID, text)
}

func handlerFreeRam(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Ошибка чтения /proc/meminfo: "+err.Error())
	}
	defer file.Close()

	var totalKB, availableKB uint64
	scanner := bufio.NewScanner(file)
	

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 2 {
			continue
		}

		val, err := strconv.ParseUint(fields[1], 10, 64)
		if err != nil {
			continue
		}

		switch fields[0] {
		case "MemTotal:":
			totalKB = val
		case "MemAvailable:":
			availableKB = val
		}

		if totalKB > 0 && availableKB > 0 {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "ошибка чтения файла:"+err.Error())
	}

	usedMB := float64(totalKB-availableKB) / 1024
	totalMB := float64(totalKB) / 1024

	text := fmt.Sprintf("RAM: %.2f MB / %.2f MB", usedMB, totalMB)
	return tgbotapi.NewMessage(msg.Chat.ID, text)
}
