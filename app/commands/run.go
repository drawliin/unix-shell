package commands

import (
	"os"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd, stdoutRedirect string) error {
	if stdoutRedirect == "" {
		cmd.Stdout = os.Stdout
		return cmd.Run()
	}

	file, err := os.Create(stdoutRedirect)
	if err != nil {
		return err
	}
	defer file.Close()

	cmd.Stdout = file
	return cmd.Run()
}
