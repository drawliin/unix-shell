package commands

import (
	"os"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd, stdRedirect stdPath) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	if stdRedirect.path == "" {
		return cmd.Run()
	}

	file, err := os.Create(stdRedirect.path)
	if err != nil {
		return err
	}
	defer file.Close()

	if stdRedirect.stdout {
		cmd.Stdout = file
		return cmd.Run()
	}
	
	if stdRedirect.stderr {
		cmd.Stderr = file
		return cmd.Run()
	}

	return cmd.Run()
}
