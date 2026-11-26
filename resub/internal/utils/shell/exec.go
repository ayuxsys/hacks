package shell

import (
	"fmt"
	"os/exec"

	"log"
)

func Exec(name string, args ...string) error {
	path, err := exec.LookPath(name)
	if err != nil {
		return fmt.Errorf("%q not present in PATH", name)
	}
	cmd := exec.Command(path, args...)
	log.Println("running commnd", cmd.String())

	return cmd.Run()
}
