package commands

import (
	"os"
	"strings"
	"unicode"
)

func appendToken(tokens *[]shellToken, current *strings.Builder, currentQuoted *bool, hadSpaceBetweenQuotes *bool) {
	if current.Len() == 0 {
		return
	}

	if *hadSpaceBetweenQuotes || len(*tokens) == 0 {
		*tokens = append(*tokens, shellToken{
			text:   current.String(),
			quoted: *currentQuoted,
		})
	} else {
		(*tokens)[len(*tokens)-1].text += current.String()
		(*tokens)[len(*tokens)-1].quoted = (*tokens)[len(*tokens)-1].quoted || *currentQuoted
	}

	current.Reset()
	*currentQuoted = false
}

func SplitTokens(input string) (string, []string, string) {
	tokens := make([]shellToken, 0)

	var current strings.Builder

	insideSingleQuote := false
	insideDoubleQuote := false
	hadSpaceBetweenQuotes := true
	currentQuoted := false
	backslash := false

	for index, rune := range input {
		if insideDoubleQuote {
			if !backslash && rune == '"' {
				currentQuoted = true
				appendToken(&tokens, &current, &currentQuoted, &hadSpaceBetweenQuotes)
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
				currentQuoted = true
				appendToken(&tokens, &current, &currentQuoted, &hadSpaceBetweenQuotes)
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
		} else if rune == '>' {
			operator := ">"
			if current.String() == "1" {
				current.Reset()
				operator = "1>"
			} else {
				appendToken(&tokens, &current, &currentQuoted, &hadSpaceBetweenQuotes)
			}
			tokens = append(tokens, shellToken{text: operator})
			hadSpaceBetweenQuotes = true
		} else if unicode.IsSpace(rune) {
			hadSpaceBetweenQuotes = true
			appendToken(&tokens, &current, &currentQuoted, &hadSpaceBetweenQuotes)
		} else {
			current.WriteRune(rune)
		}
	}

	// Last field might end at EOF.
	appendToken(&tokens, &current, &currentQuoted, &hadSpaceBetweenQuotes)

	redirectPath := ""
	filteredArgs := make([]string, 0, len(tokens))

	for i := 0; i < len(tokens); i++ {
		if (tokens[i].text == ">" || tokens[i].text == "1>") && !tokens[i].quoted {
			if i+1 < len(tokens) {
				redirectPath = tokens[i+1].text
				i++
			}
			continue
		}

		filteredArgs = append(filteredArgs, tokens[i].text)
	}

	if len(filteredArgs) == 0 {
		return "", nil, redirectPath
	}

	return filteredArgs[0], filteredArgs[1:], redirectPath
}

func winToUnixPath(homeDir string) string {
	homeDir = strings.ReplaceAll(homeDir, "\\", "/")

	if len(homeDir) >= 2 && homeDir[1] == ':' {
		drive := strings.ToLower(homeDir[:1])
		homeDir = "/" + drive + homeDir[2:]
	}

	return homeDir
}
