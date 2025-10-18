package telegram

import (
	"github.com/ashkanamani/madkings/internal/service"
	"github.com/ashkanamani/madkings/internal/telegram/teleprompt"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"time"
)

type Telegram struct {
	App        *service.App
	bot        *telebot.Bot
	TelePrompt *teleprompt.TelePrompt
}

func NewTelegram(app *service.App, token string) (*Telegram, error) {
	telegram := &Telegram{App: app, TelePrompt: teleprompt.NewTelePrompt()}

	pref := telebot.Settings{
		Token:   token,
		Poller:  &telebot.LongPoller{Timeout: 60 * time.Second},
		OnError: telegram.OnError,
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Errorln("could not connect to telegram servers")
		return nil, err
	}
	telegram.bot = bot
	telegram.setupHandlers()

	return telegram, nil
}

func (t *Telegram) Start() {
	logrus.Infoln("starting telegram bot")
	t.bot.Start()
}
