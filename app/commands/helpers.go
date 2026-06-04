package commands

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
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

func ArgsParser(input string) (string, []string) {
	var cmdArgs []string
	// get command name
	cmdName := strings.Fields(input)[0]

	// remove command from input and trim the rest
	input = strings.TrimSpace(strings.TrimPrefix(input, cmdName))
	// remove adjacent quotes
	input = strings.ReplaceAll(input, "''", "")

	countQuote := strings.Count(input, "'")
	if countQuote == 0 {
		return cmdName, sanitizeArgs(strings.Fields(input))
	}

	var arg strings.Builder
	foundQuote := false

	// get args
	for _, c := range input {
		switch c {
		case '\'':
			arg.WriteRune(c)
			if foundQuote {
				cmdArgs = append(cmdArgs, arg.String())
				arg.Reset()
				foundQuote = false
			} else {
				foundQuote = true
			}

		case ' ':
			if foundQuote {
				arg.WriteRune(c)
			}

		default:
			arg.WriteRune(c)
		}
	}

	// add any remaining args
	if arg.Len() > 0 {
		cmdArgs = append(cmdArgs, arg.String())
		arg.Reset()
	}

	return cmdName, sanitizeArgs(cmdArgs)
}

// Sanitize args
func sanitizeArgs(cmdArgs []string) []string {
	for i := 0; i < len(cmdArgs); i++ {
		if strings.HasPrefix(cmdArgs[i], "'") && strings.HasSuffix(cmdArgs[i], "'") {
			cmdArgs[i] = strings.TrimPrefix(strings.TrimSuffix(cmdArgs[i], "'"), "'")
		} else {
			cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "~", os.Getenv("HOME"))
		}
	}

	return cmdArgs
}
