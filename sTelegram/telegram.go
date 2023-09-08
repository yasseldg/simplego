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
	Bot    *tgbotapi.BotAPI
	ChatId int64
	Token  string
	Func   ReadFunc
}

type SendObject struct {
	ChatId  int64  `json:"chat_id"`
	Message string `json:"msg"`
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
		read_func = defaultFunc
	}
	sLog.Debug("NewTelegramBot: chat_id: %d, read_func: %v", chat_id, read_func)
	return &Bot{ChatId: chat_id, Token: token, Func: read_func}
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

	t.Send(SendObject{ChatId: t.ChatId, Message: fmt.Sprintf("Starting ... %s ", t.Bot.Self.UserName)})

	go t.read()
}

func (t Bot) Send(obj SendObject) {
	if t.Bot == nil {
		sLog.Error("TelegramBot.Send: bot is nil")
		return
	}

	if obj.ChatId == 0 {
		obj.ChatId = t.ChatId
	}

	newMsg := tgbotapi.NewMessage(obj.ChatId, obj.Message)
	_, err := t.Bot.Send(newMsg)
	if err != nil {
		sLog.Error("TelegramBot.Send: %s", err)
		t.Send(SendObject{ChatId: t.ChatId, Message: fmt.Sprintf("Error sending ... %s \n Chat ID:    %d \n Message:  %s", err, obj.ChatId, obj.Message)})
	}
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
		msg := t.Func(&update)
		if len(msg) > 0 {
			t.Send(SendObject{ChatId: t.ChatId, Message: msg})
		}
	}
}

func defaultFunc(update *tgbotapi.Update) string {
	if update.Message == nil {
		return ""
	}
	return fmt.Sprintf("%s: %s", update.Message.From.UserName, update.Message.Text)
}
