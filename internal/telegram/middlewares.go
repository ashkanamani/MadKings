package telegram

import (
	"context"
	"github.com/ashkanamani/madkings/internal/entity"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) registerMiddleWare(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(c telebot.Context) error {
		acc := entity.Account{
			ID:        c.Sender().ID,
			FirstName: c.Sender().FirstName,
			Username:  c.Sender().Username,
		}
		_, _, err := t.App.Account.CreateOrUpdate(context.Background(), acc)
		if err != nil {
			return err
		}
		return next(c)
	}
}
