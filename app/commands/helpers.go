package commands

import (
	"os"
	"strings"
	"unicode"
)

func SplitTokens(input string) (string, []string) {
	args := make([]string, 0)

	var current strings.Builder

	insideSingleQuote := false
	insideDoubleQuote := false
	hadSpaceBetweenQuotes := true
	backslash := false

	for index, rune := range input {
		if insideDoubleQuote {
			if !backslash && rune == '"' {
				if hadSpaceBetweenQuotes {
					args = append(args, current.String())
				} else {
					// just concatenate to previous string
					args[len(args)-1] += current.String()
				}
				current.Reset()
				insideDoubleQuote = false
				hadSpaceBetweenQuotes = false
			} else if !backslash && rune == '\\' {
				backslash = true
			} else if backslash {
				switch rune {
				case '\\', '"', '$', '\n':
					backslash = false
					current.WriteRune(rune)
				default:
					backslash = false
					current.WriteRune('\\')
					current.WriteRune(rune)
				}
			} else {
				current.WriteRune(rune)
			}
		} else if insideSingleQuote {
			if rune == '\'' {
				if hadSpaceBetweenQuotes {
					args = append(args, current.String())
				} else {
					// just concatenate to previous string
					args[len(args)-1] += current.String()
				}
				current.Reset()
				insideSingleQuote = false
				hadSpaceBetweenQuotes = false
			} else {
				current.WriteRune(rune)
			}
		} else if backslash {
			backslash = false
			current.WriteRune(rune)
		} else if rune == '\\' {
			backslash = true
		} else if rune == '\'' {
			insideSingleQuote = true
		} else if rune == '"' {
			insideDoubleQuote = true
		} else if rune == '~' {
			if current.Len() > 0 || (current.Len() == 0 && index < len(input)-1 && input[index+1] != '/') {
				current.WriteRune(rune)
			} else {
				homeDir := winToUnixPath(os.Getenv("HOME"))
				for _, c := range homeDir {
					current.WriteRune(c)
				}
			}
		} else if unicode.IsSpace(rune) {
			hadSpaceBetweenQuotes = true
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		} else {
			current.WriteRune(rune)
		}
	}

	// Last field might end at EOF.
	if current.Len() > 0 {
		if hadSpaceBetweenQuotes {
			args = append(args, current.String())
		} else {
			// just concatenate to previous string
			args[len(args)-1] += current.String()
		}
	}

	return args[0], args[1:]
}

func winToUnixPath(homeDir string) string {
	homeDir = strings.ReplaceAll(homeDir, "\\", "/")

	if len(homeDir) >= 2 && homeDir[1] == ':' {
		drive := strings.ToLower(homeDir[:1])
		homeDir = "/" + drive + homeDir[2:]
	}

	return homeDir
}
