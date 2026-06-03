package commands

import (
	"errors"
	"os"
	"path/filepath"
)

func FindExecutable(pathenv string, cmd string) (string, error) {
	if pathenv == "" {
		return "", errors.New("No PATH Provided")
	}

	var err error

	paths := filepath.SplitList(pathenv)

	for _, p := range paths {
		full := filepath.Join(p, cmd)
		if isExecutable(full) {
			return full, err
		}
	}

	return "", errors.New("Not an executable")
}

func isExecutable(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}

	if info.IsDir() {
		return false
	}

	return info.Mode()&0111 != 0
}
