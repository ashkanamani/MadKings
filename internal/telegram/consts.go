package telegram

import "time"

var (
	DefaultInputTimeout     = time.Minute * 5
	DefaultInputTimeoutText = "We were waiting for you but you did not send anything. Please send message when you come ⏰"
)
