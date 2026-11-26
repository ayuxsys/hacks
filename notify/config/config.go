package config

import (
	"os"
	"path/filepath"

	"github.com/ayuxsys/notify/must"
)

var (
	BaseDir     = filepath.Join(must.Eval(os.UserConfigDir()), "notify")
	DefaultPath = BaseDir + "/config.yaml"
)

type Config struct {
	Notify NotifyCfg `yaml:"notify"`
}

type NotifyCfg struct {
	Telegram TelegramCfg `yaml:"telegram"`
}

type TelegramCfg struct {
	Enabled  bool   `yaml:"enabled"`
	BotToken string `yaml:"bot_token"`
	ChatID   int64  `yaml:"chat_id"`
}
