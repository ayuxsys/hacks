package main

import (
	"log"
	"os"

	"github.com/fatih/color"
)

var (
	infoL = log.New(os.Stderr, color.CyanString("[info] "), 0)
	warnL = log.New(os.Stderr, color.YellowString("[warn] "), 0)
	errL  = log.New(os.Stderr, color.RedString("[error] "), log.Lshortfile)
	okL   = log.New(os.Stderr, color.GreenString("[ok] "), 0)
)
