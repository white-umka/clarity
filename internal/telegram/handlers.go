package telegram

import (
	"bufio"
	"clarity/internal/ascii"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

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

func handleFreeRam(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.MessageConfig {
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

func handleASCII(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) tgbotapi.Chattable {
	// Проверяем, прислал ли пользователь фото
	if len(msg.Photo) == 0 {
		return tgbotapi.NewMessage(msg.Chat.ID, "Пожалуйста, отправьте фотографию.")
	}

	// Берем фото в наилучшем качестве (последний элемент среза)
	photo := msg.Photo[len(msg.Photo)-1]

	// Получаем прямую ссылку на скачивание файла с серверов Telegram
	fileURL, err := bot.GetFileDirectURL(photo.FileID)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Не удалось получить ссылку на фото.")
	}

	// 1. Скачиваем картинку
	img, err := ascii.GetImage(fileURL)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Ошибка при загрузке изображения.")
	}

	// 2. Переводим в матрицу символов
	art := ascii.Converter(img)

	// 3. Рисуем в PNG в памяти
	pngFile, err := ascii.Render(art)
	if err != nil {
		return tgbotapi.NewMessage(msg.Chat.ID, "Ошибка при генерации картинки.")
	}

	// Возвращаем готовое фото для отправки через bot.Send()
	return tgbotapi.NewPhoto(msg.Chat.ID, pngFile)
}
