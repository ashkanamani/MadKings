package telegram

import (
	"github.com/ashkanamani/madkings/internal/service"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
	"time"
)

type Telegram struct {
	App *service.App
	bot *telebot.Bot
}

func NewTelegram(app *service.App, token string) (*Telegram, error) {
	pref := telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 60 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		logrus.WithError(err).Errorln("could not connect to telegram servers")
		return nil, err
	}
	telegram := &Telegram{App: app, bot: b}
	telegram.setupHandlers()

	return telegram, nil
}

func (t *Telegram) Start() {
	logrus.Infoln("starting telegram bot")
	t.bot.Start()
}
