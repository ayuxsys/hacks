package main

import (
	"io"
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	infoL = log.New(os.Stderr, color.CyanString("[info] "), 0)
	warnL = log.New(os.Stderr, color.YellowString("[warn] "), 0)
	errL  = log.New(os.Stderr, color.RedString("[error]"), 0)
	okL   = log.New(os.Stderr, color.GreenString("[ok] "), 0)
	opL   = log.New(os.Stdout, "", 0)
)

func setLogOp(file *os.File) {
	opL.SetOutput(io.MultiWriter(file, os.Stdout))
}
