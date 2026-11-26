package log

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	Info  = log.New(os.Stderr, color.CyanString("[info] "), log.Ltime)
	Error = log.New(os.Stderr, color.RedString("[error] "), log.Ltime)
)
