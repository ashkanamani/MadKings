package telegram

import (
	"fmt"
	"gopkg.in/telebot.v4"
)

func (t *Telegram) start(c telebot.Context) error {
	isJustCreated := c.Get("is_just_created").(bool)
	if !isJustCreated {
		// TODO: ..
		
	}
	msg, err := t.Input(c, InputConfig{
		Prompt:    "Hi, What's your name?",
		OnTimeout: "Reached timeout!",
	})
	if err != nil {
		return err
	}
	return c.Reply(fmt.Sprintf("Your name is %s.", msg.Text))
}
