package log

import (
	"log"
	"os"
)

var (
	Info = log.New(os.Stderr, "[info] ", log.Lshortfile)
)
