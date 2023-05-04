package sTelegram

import (
	"fmt"

	"github.com/yasseldg/simplego/sConv"
	"github.com/yasseldg/simplego/sEnv"
	"github.com/yasseldg/simplego/sLog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Environment Variables
// ** TelegramDebug bool		# by default false
// ** TelegramBotToken string
// ** TelegramChatID int64

type Bot struct {
	Bot      *tgbotapi.BotAPI
	ChatId   int64
	Token    string
	ReadFunc ReadFunc
}

type ReadFunc func(update *tgbotapi.Update) string

func NewBot(token string, chat_id int64, read_func ReadFunc) *Bot {
	if chat_id == 0 {
		chat_id = sConv.GetInt64(sEnv.Get("TelegramChatID", ""))
	}
	if len(token) == 0 {
		token = sEnv.Get("TelegramBotToken", "")
	}
	if read_func == nil {
		read_func = defaultRead
	}
	sLog.Debug("NewTelegramBot: chat_id: %d, read_func: %v", chat_id, read_func)
	return &Bot{ChatId: chat_id, Token: token, ReadFunc: read_func}
}

func defaultRead(update *tgbotapi.Update) string {
	if update.Message == nil {
		return ""
	}
	return fmt.Sprintf("%s: %s", update.Message.From.UserName, update.Message.Text)
}

func (t *Bot) Start() {
	sLog.Info("Start Telegram BOT ...")

	bot, err := tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		sLog.Error("TelegramBot.Start: %s", err.Error())
		return
	}
	bot.Debug = sConv.GetBool(sEnv.Get("TelegramDebug", "false"))
	t.Bot = bot

	go t.read()
}

func (t *Bot) read() {
	if t.Bot == nil {
		sLog.Error("TelegramBot.read: bot is nil")
		return
	}

	sLog.Debug("Authorized on account %s", t.Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := t.Bot.GetUpdatesChan(u)
	for update := range updates {
		str := t.ReadFunc(&update)
		if len(str) > 0 {
			t.Send(str)
		}
	}
}

func (t *Bot) Send(msg string) {
	if t.Bot == nil {
		sLog.Error("TelegramBot.Send: bot is nil")
		return
	}

	newMsg := tgbotapi.NewMessage(t.ChatId, msg)
	_, err := t.Bot.Send(newMsg)
	if err != nil {
		sLog.Error("TelegramBot.Send: %s", err)
	}
}
