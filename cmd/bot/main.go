package main

import (
	"clarity/internal/telegram"
)

func main() {
	bot := telegram.NewBot()
	telegram.StartUpdates(bot)
}

