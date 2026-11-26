package notifier

import "github.com/ayuxsys/hacks/notify/config"

type Notify struct {
	Cfg  config.Config
	File string
	Msg  string
}

func New(cfg config.Config) *Notify {
	return &Notify{
		Cfg: cfg,
	}
}
