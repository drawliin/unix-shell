package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var SHELL_COMMANDS = []string{"exit", "echo", "type"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, _ := reader.ReadString('\n')

		// check exit command to exit the program
		if isExitCommand(command) {
			break
		}

		// check echo command
		if isEchoCommand(command) {
			args := strings.TrimSpace(command[4:])
			fmt.Println(args)
			continue
		}

		// check type command
		if isTypeCommand(command) {
			rawArgs := strings.TrimSpace(command[4:])
			args := strings.Split(rawArgs, " ")

			for _, arg := range args {
				if slices.Contains(SHELL_COMMANDS, arg) {
					fmt.Printf("%s is a shell builtin\n", arg)
				} else {
					fmt.Printf("%s: not found\n", arg)
				}
			}

			continue
		}

		fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
	}
}

func isExitCommand(s string) bool {
	if strings.TrimSpace(s) == "exit" {
		return true
	}
	return false
}

func isEchoCommand(s string) bool {
	str := strings.TrimSpace(s)
	if strings.HasPrefix(str, "echo") {
		args := strings.Split(str, " ")
		if args[0] == "echo" {
			return true
		}
	}
	return false
}

func isTypeCommand(s string) bool {
	str := strings.TrimSpace(s)
	if strings.HasPrefix(str, "type") {
		args := strings.Split(str, " ")
		if args[0] == "type" {
			return true
		}
	}
	return false
}
