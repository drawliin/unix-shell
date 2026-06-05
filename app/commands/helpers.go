package commands

import (
	"os"
	"strings"
)

func ArgsParser(input string) (string, []string) {
	var cmdArgs []string
	// get command name
	cmdName := strings.Fields(input)[0]

	// remove command from input and trim the rest
	input = strings.TrimSpace(strings.TrimPrefix(input, cmdName))

	// remove adjacent quotes
	input = strings.ReplaceAll(input, "''", "")
	input = strings.ReplaceAll(input, "\"\"", "")

	if strings.Count(input, "'") == 0 && strings.Count(input, "\"") == 0 {
		return cmdName, sanitizeArgs(strings.Fields(input))
	}

	var arg strings.Builder

	// cast to runes
	inputRunes := []rune(input)

	// get args
	for i := 0; i < len(inputRunes); {
		switch inputRunes[i] {
		case '\'':
			index := appendUntilEnd(inputRunes, "'", i)
			for i <= index {
				arg.WriteRune(inputRunes[i])
				i++
			}
			cmdArgs = append(cmdArgs, arg.String())
			arg.Reset()
			i--

		case '"':
			index := appendUntilEnd(inputRunes, "\"", i)
			for i <= index {
				arg.WriteRune(inputRunes[i])
				i++
			}
			cmdArgs = append(cmdArgs, arg.String())
			arg.Reset()
			i--

		case ' ':
			i++
			continue

		default:
			arg.WriteRune(inputRunes[i])
		}

		i++
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
		length := len(cmdArgs[i])
		if cmdArgs[i][0] == '\'' && cmdArgs[i][length-1] == '\'' {
			cmdArgs[i] = cmdArgs[i][1 : length-1]
		} else if cmdArgs[i][0] == '"' && cmdArgs[i][length-1] == '"' {
			cmdArgs[i] = cmdArgs[i][1 : length-1]
		} else {
			cmdArgs[i] = strings.ReplaceAll(cmdArgs[i], "~", os.Getenv("HOME"))
		}
	}

	return cmdArgs
}

func appendUntilEnd(inputRunes []rune, param string, index int) int {
	for i := index + 1; i < len(inputRunes); i++ {
		if string(inputRunes[i]) == param && i < len(inputRunes)-1 && inputRunes[i+1] != ' ' {
			param = " "
		} else if string(inputRunes[i]) == param {
			return i
		} else if param == " " && i == len(inputRunes)-1 {
			return i
		}
	}

	return -1
}
