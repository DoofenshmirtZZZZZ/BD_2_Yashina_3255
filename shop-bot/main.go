package main

import (
	"fmt"
	tgbot "shop-bot/bot"
	"shop-bot/database"
)

func main() {
	fmt.Println("start")
	err := database.CheckTable()

	if err == nil {
		tgbot.Start()
	}

	fmt.Println("stop")
}
