package notifier

import (
	"fmt"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (n *Notify) Telegram() error {
	cfg := n.Cfg.Notify.Telegram

	if !cfg.Enabled {
		return nil
	}

	bot, err := tgbotapi.NewBotAPI(cfg.BotToken)
	if err != nil {
		return fmt.Errorf("create telegram bot: %w", err)
	}

	if n.Msg != "" {
		_, err := bot.Send(tgbotapi.NewMessage(cfg.ChatID, n.Msg))
		if err != nil {
			return fmt.Errorf("send message 'bot.Send': %w", err)
		}
	}

	if n.File != "" {
		data, err := os.ReadFile(n.File)
		if err != nil {
			return fmt.Errorf("read msg file 'os.ReadFile': %w", err)
		}
		doc := tgbotapi.NewDocument(cfg.ChatID,
			tgbotapi.FileBytes{
				Name:  filepath.Base(n.File),
				Bytes: data,
			})
		_, err = bot.Send(doc)
		if err != nil {
			return fmt.Errorf("send file 'bot.Send': %w", err)
		}
	}

	return nil
}
