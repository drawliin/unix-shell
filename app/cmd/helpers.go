package cmd

import "strings"

func IsExitCommand(s string) bool {
	if strings.TrimSpace(s) == "exit" {
		return true
	}
	return false
}

func IsEchoCommand(s string) bool {
	str := strings.TrimSpace(s)
	if strings.HasPrefix(str, "echo") {
		args := strings.Split(str, " ")
		if args[0] == "echo" {
			return true
		}
	}
	return false
}

func IsTypeCommand(s string) bool {
	str := strings.TrimSpace(s)
	if strings.HasPrefix(str, "type") {
		args := strings.Split(str, " ")
		if args[0] == "type" {
			return true
		}
	}
	return false
}