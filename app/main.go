package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"shell/app/cmd"
)

var SHELL_COMMANDS = []string{"exit", "echo", "type"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, _ := reader.ReadString('\n')

		// check exit command to exit the program
		if cmd.IsExitCommand(command) {
			break
		}

		// check echo command
		if cmd.IsEchoCommand(command) {
			args := strings.TrimSpace(command[4:])
			fmt.Println(args)
			continue
		}

		// check type command
		if cmd.IsTypeCommand(command) {
			rawArgs := strings.TrimSpace(command[4:])
			args := strings.Split(rawArgs, " ")

			for _, arg := range args {
				if slices.Contains(SHELL_COMMANDS, arg) {
					fmt.Printf("%s is a shell builtin\n", arg)
					break
				}
				if exe, err := cmd.FindExecutable(os.Getenv("PATH"), arg); err == nil {
					fmt.Printf("%s is %s\n", arg, exe)
				} else {
					fmt.Printf("%s: not found\n", arg)
				}
			}

			continue
		}

		fmt.Printf("%s: command not found\n", strings.TrimSpace(command))
	}
}

