package commands

import (
	"io"
	"os"
	"os/exec"
)

func RunCommand(cmd *exec.Cmd, stdRedirect stdPath) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if stdRedirect.path == "" {
		return cmd.Run()
	}

	file, err := os.OpenFile(stdRedirect.path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Use a normal writable handle for child processes on every platform.
	// On Windows/MSYS, inheriting an O_APPEND-style stdout can make tools like
	// cat/ls fail with "Bad file descriptor", so we seek to EOF ourselves.
	if stdRedirect.appendEOF {
		file.Seek(0, io.SeekEnd)
	} else {
		// Plain `>` should replace the file contents before the child runs.
		file.Truncate(0)
		file.Seek(0, 0)
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
