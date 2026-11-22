package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
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
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	infoL.Printf("running command: %q", cmd.String())
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("cmd=%q, stderr=%s, error=%v", cmd.String(), stderr.String(), err)
	}

	return stdout.String(), nil
}
