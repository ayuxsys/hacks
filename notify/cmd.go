package main

import (
	"github.com/ayuxsys/hacks/notify/config"
	"github.com/spf13/pflag"
)

type cmdi struct {
	notify  notifyOpts
	cfgFile string
}

type notifyOpts struct {
	file string
	msg  string
}

func (c *cmdi) parseFlags() {
	pflag.StringVarP(&c.notify.file, "file", "f", "", "path to the file to send as msg")
	pflag.StringVarP(&c.cfgFile, "config", "c", config.DefaultPath, "path to the config file to use")
	pflag.StringVarP(&c.notify.msg, "msg", "m", "", "short msg to send (len < 4096)")
	pflag.Usage = func() {
		println("Usage: notify [flags]")
		pflag.PrintDefaults()
	}
	pflag.Parse()
}
