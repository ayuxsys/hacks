package log

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
	"golang.org/x/term"
)

// loggers
var (
	Info   = log.New(os.Stderr, "[info] ", log.Lshortfile|log.Ltime)
	Warn   = log.New(os.Stderr, "[warn] ", log.Lshortfile|log.Ltime)
	Err    = log.New(os.Stderr, "[error]", log.Lshortfile|log.Ltime)
	Output = log.New(os.Stdout, "", 0) // output logger
)

func init() {
	if term.IsTerminal(int(os.Stderr.Fd())) {
		Info.SetPrefix(color.BlueString("[info] "))
		Warn.SetPrefix(color.YellowString("[warn] "))
		Err.SetPrefix(color.RedString("[error] "))
	}
}

func SetOutput(out *os.File) {
	multiWriter := io.MultiWriter(os.Stdout, out)
	Output.SetOutput(multiWriter)
}
