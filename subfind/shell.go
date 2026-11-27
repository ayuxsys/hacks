package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

var (
	stderrLogFile string
)

type shell struct {
	name  string
	stdin string
}

// execCmd executes a command with it's args and retunrs stdout or err which also contains stderr
func (s shell) execute(args ...string) (string, error) {
	cmd := exec.Command(s.name, args...)
	cmd.Stdin = strings.NewReader(s.stdin)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout

	var stderr bytes.Buffer
	if stderrLogFile != "" {
		file := must(os.OpenFile(stderrLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644))
		defer file.Close()
		cmd.Stderr = io.MultiWriter(file, &stderr)
	} else {
		cmd.Stderr = &stderr
	}

	s.logCmd(cmd.String())
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cmd=%q, stderr=%s, error=%v", cmd.String(), stderr.String(), err)
	}

	return stdout.String(), nil
}

func (s shell) logCmd(st string) {
	switch s.stdin {
	case "":
		infoL.Printf("running command: %q", st)
	default:
		infoL.Printf("running command: %q, domain(s)=%d", st, len(strings.Split(s.stdin, "\n")))
	}
}
