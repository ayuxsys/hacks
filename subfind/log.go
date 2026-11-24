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
)

var (
	opL  = log.New(os.Stdout, "", 0)
	opAL = log.New(io.Discard, "", 0)
)

func setLogOp(path string) func() {
	file, err := os.OpenFile(path, os.O_TRUNC|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	opL.SetOutput(io.MultiWriter(file, os.Stdout))
	return func() { file.Close() }
}

func setLogAllOP(path string) func() {
	file, err := os.OpenFile(path, os.O_TRUNC|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	opAL.SetOutput(file)
	return func() { defer file.Close() }
}
