package telegram

import (
	"context"
	"errors"
	"fmt"
	"github.com/ashkanamani/madkings/internal/entity"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) registerMiddleWare(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			Username:  c.Sender().Username,
		}
		acc, created, err := t.App.Account.CreateOrUpdate(context.Background(), acc)
		if err != nil {
			return err
		}

		c.Set("account", acc)
		c.Set("is_just_created", created)

		return next(c)
	}
}

func (t *Telegram) OnError(err error, c telebot.Context) {
	if errors.Is(err, ErrorInputTimeout) {
		return
	}
	errId := uuid.New().String()

	logrus.WithError(err).WithField("tracing_id", errId).Errorln("unhandled error")
	c.Reply(fmt.Sprintf("❌There is a problem while processing your message"))
}
