package main

import (
	"io"
	"os"

	"github.com/ayuxsys/notify/config"
	"github.com/ayuxsys/notify/log"
	"github.com/ayuxsys/notify/must"
	"github.com/ayuxsys/notify/notifier"
)

func init() {
	config.Setup()
}

func main() {
	c := &cmdi{}
	c.parseFlags()
	c.Run()
}

func (c *cmdi) Run() {
	notify := notifier.New(config.Load(c.cfgFile))
	notify.File = c.notify.file
	notify.Msg = c.deferMsg()

	if err := notify.Telegram(); err != nil {
		log.Error.Fatal(err)
	}
}

// deferMsg determines which input to use as the message for deferMsg.
// If input is piped and "-m" is an empty string, it reads from stdin
// and uses the piped input as the message.
func (c *cmdi) deferMsg() string {
	if c.notify.msg != "" {
		return c.notify.msg
	}
	if isPipedStdin() {
		return string(must.Eval(io.ReadAll(os.Stdin)))
	}
	return "unknown"
}

func isPipedStdin() bool {
	fi := must.Eval(os.Stdin.Stat())
	return (fi.Mode() & os.ModeCharDevice) == 0
}
