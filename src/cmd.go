package AdsBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func CloseKeyboard(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	bot.Send(msg)
}

func StartDialog(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, "+update.Message.From.FirstName)
	msg.ReplyMarkup = StartKey()
	bot.Send(msg)
}

func CancelDialog(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Отмена")
	msg.ReplyMarkup = StartKey()
	bot.Send(msg)
}

func AboutBot(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Шикарный бот")
	bot.Send(msg)
}

func GetAds(bot *tgbotapi.BotAPI, update *tgbotapi.Update, off int) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сегодня "+CountAds()+" объявлений")
	msg.ReplyMarkup = AddKeyAds(off)
	bot.Send(msg)
}

func AddUserToDB(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	resinsert := InsertUser(update.Message.From.ID, update.Message.From.FirstName)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, resinsert)
	msg.ReplyMarkup = StartKey()
	bot.Send(msg)
}

func EditAdsButtons(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {

}
