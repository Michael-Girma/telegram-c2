package models

import "time"

type Agent struct {
	TelegramID  string
	LastPing    time.Time
	LastCommand string
}
