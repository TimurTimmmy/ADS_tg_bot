package AdsBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func AddKeyAds(off int) tgbotapi.InlineKeyboardMarkup {
	if off < 0 {
		off = 0
	}
	forkeyads := DBRequest(off)
	keyboard := tgbotapi.InlineKeyboardMarkup{}

	for _, r := range forkeyads {
		var row1 []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(r.adsName+" Цена = "+strconv.Itoa(r.adsPrice), r.adsName+" Цена = "+strconv.Itoa(r.adsPrice))
		row1 = append(row1, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row1)
	}

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, NaviKeys())
	return keyboard
}

func NextPrevKeys(off int, update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	if off < 0 {
		off = 0
	}
	forkeyads := DBRequest(off)

	keyboard := tgbotapi.InlineKeyboardMarkup{}
	for _, r := range forkeyads {
		var row1 []tgbotapi.InlineKeyboardButton
		btn := tgbotapi.NewInlineKeyboardButtonData(r.adsName+" Цена = "+strconv.Itoa(r.adsPrice), r.adsName+" Цена = "+strconv.Itoa(r.adsPrice))
		row1 = append(row1, btn)
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row1)
	}
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, NaviKeys())

	m := tgbotapi.NewEditMessageReplyMarkup(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, keyboard)
	bot.Send(m)
}

func NaviKeys() (navirow []tgbotapi.InlineKeyboardButton) {
	navirow = append(navirow, tgbotapi.NewInlineKeyboardButtonData(prev, prev))
	navirow = append(navirow, tgbotapi.NewInlineKeyboardButtonData(next, next))
	return navirow
}

func StartKey() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Просмотр"),
			tgbotapi.NewKeyboardButton("Подать"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Профиль"),
			tgbotapi.NewKeyboardButton(about),
		),
	)
	return keyboard
}

func StartNew(bot *tgbotapi.BotAPI, update *tgbotapi.Update) tgbotapi.InlineKeyboardMarkup {
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	row := make([]tgbotapi.InlineKeyboardButton, 1)
	row = append(row, tgbotapi.NewInlineKeyboardButtonData("Про бота", about))
	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, row)
	return keyboard
}
