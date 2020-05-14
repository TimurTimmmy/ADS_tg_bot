package AdsBot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func SkeyBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	var (
		fromID  int
		editvar int
		offset  int
	)

	for update := range updates {

		if update.CallbackQuery != nil {
			data := update.CallbackQuery.Data
			if data == "Изменить Имя" {
				editvar = 1
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Отмена")))
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите новое имя")
				msg.ReplyMarkup = keyboard
				fromID = update.CallbackQuery.From.ID
				bot.Send(msg)
				continue
			}
			if data == "Изменить Адрес" {
				editvar = 2
				keyboard := tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Отмена")))
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите новый адрес")
				msg.ReplyMarkup = keyboard
				fromID = update.CallbackQuery.From.ID
				bot.Send(msg)
				continue
			}

			if data == next {
				offset += 5
				NextPrevKeys(offset, &update, bot)
				continue
			}
		}

		if update.Message != nil && fromID != 0 && editvar != 0 {
			if update.Message.Text == "Отмена" {
				fromID, editvar = 0, 0
				CancelDialog(bot, &update)
				continue
			}
			if editvar == 1 {
				UpdateUserName(fromID, update.Message.Text)
			} else if editvar == 2 {
				UpdateUserAddress(fromID, update.Message.Text)
			}
			fromID, editvar = 0, 0
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Готово")
			msg.ReplyMarkup = StartKey()
			bot.Send(msg)
		}

		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				StartDialog(bot, &update)
				continue

			case "close":
				CloseKeyboard(bot, &update)
				continue
			}

			switch update.Message.Text {
			case "Просмотр":
				offset = 0
				GetAds(bot, &update, offset)
				continue

			case "Профиль":
				chatuser := update.Message.From.ID
				if SeachUser(chatuser) != 0 {
					user := FindUser(chatuser)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ваш профиль: \n \n"+"Имя: "+user.userName+"\n"+"Адрес: "+user.userAddress)

					keyboard := tgbotapi.InlineKeyboardMarkup{}
					var row []tgbotapi.InlineKeyboardButton
					btn := tgbotapi.NewInlineKeyboardButtonData("Изменить Имя", "Изменить Имя")
					btn1 := tgbotapi.NewInlineKeyboardButtonData("Изменить Адрес", "Изменить Адрес")
					row = append(row, btn)
					row = append(row, btn1)
					keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пользователь не найден")
					keyboard := tgbotapi.NewReplyKeyboard(
						tgbotapi.NewKeyboardButtonRow(
							tgbotapi.NewKeyboardButton("Добавить"),
							tgbotapi.NewKeyboardButton("Отмена"),
						),
					)
					msg.ReplyMarkup = keyboard
					bot.Send(msg)
				}
			case "Добавить":
				AddUserToDB(bot, &update)

			case "Отмена":
				StartDialog(bot, &update)
				continue
			}
		}
	}
}

func SeachUser(id int) (result int) {
	users := GetUsers()
	for _, r := range users {
		if id == r.userId {
			result = id
		} else {
			result = 0
		}
	}
	return result
}
