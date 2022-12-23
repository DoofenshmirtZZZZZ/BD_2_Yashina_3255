package tgbot

import (
	"fmt"
	"os"
	"shop-bot/database"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Авторизация %s", bot.Self.UserName)

	temp := tgbotapi.NewUpdate(0)
	temp.Timeout = 60

	updates := bot.GetUpdatesChan(temp)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Command() {
			case "Добавить товар":
				_, textProduct, _ := strings.Cut(update.Message.Text, update.Message.Command()+" ")

				if len(textProduct) == 0 {
					messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не ввели товар.")
					bot.Send(messTxt)
					continue
				}

				err := database.Set_product(update.Message.From.ID, textProduct)
				message := ""

				if err != nil {
					message = "Ошибка"
				} else {
					message = "Готово"
				}

				messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(messTxt)
				continue
			case "delete":
				_, textProduct, _ := strings.Cut(update.Message.Text, update.Message.Command()+" ")

				if len(textProduct) == 0 {
					messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы не ввели товар.")
					bot.Send(messTxt)
					continue
				}

				err := database.Delete_product(update.Message.From.ID, textProduct)

				message := ""

				if err != nil {
					message = "Ошибка"
				} else {
					message = "Готово"
				}

				messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(messTxt)
				continue
			case "list":
				goods, err := database.Get_list_product(update.Message.From.ID)

				message := ""

				if err != nil {
					messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка")
					bot.Send(messTxt)
					continue
				}

				if len(goods) > 0 {

					message = "Список товаров:\n"

					for i := 0; i < len(goods); i++ {
						message += strconv.Itoa(i+1) + ") " + goods[i] + "\n"
					}
				} else {
					message = "Список товаров пуст."
				}

				messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(messTxt)
				continue
			case "help":
				message := "/add - Добавить товар.\n/delete - Удалить товар.\n/list - Список товаров.\n/help - Команды."
				messTxt := tgbotapi.NewMessage(update.Message.Chat.ID, message)
				bot.Send(messTxt)
				continue
			default:
				bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Упс, не корректный запрос..."))
			}
		}
	}
}
