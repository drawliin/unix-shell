package commands

import "strings"

func IsExitCommand(s string) bool {
	if s == "exit" {
		return true
	}
	return false
}

func IsEchoCommand(s string) bool {
	if strings.HasPrefix(s, "echo") {
		args := strings.Split(s, " ")
		if args[0] == "echo" {
			return true
		}
	}
	return false
}

func IsTypeCommand(s string) bool {
	if strings.HasPrefix(s, "type") {
		args := strings.Split(s, " ")
		if args[0] == "type" {
			return true
		}
	}
	return false
}
