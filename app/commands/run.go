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

	flags := os.O_CREATE | os.O_WRONLY

	if stdRedirect.appendEOF {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	file, err := os.OpenFile(stdRedirect.path, flags, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	if stdRedirect.stdout {
		cmd.Stdout = file
	}

	if stdRedirect.stderr {
		cmd.Stderr = file
	}

	return cmd.Run()
}
