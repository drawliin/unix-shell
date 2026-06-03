package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"

	"shell/app/commands"
)

var SHELL_COMMANDS = []string{"exit", "echo", "type"}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// check exit command to exit the program
		if commands.IsExitCommand(input) {
			break
		}

		// check echo command
		if commands.IsEchoCommand(input) {
			args := strings.TrimSpace(input[4:])
			fmt.Println(args)
			continue
		}

		// check type command
		if commands.IsTypeCommand(input) {
			rawArgs := strings.TrimSpace(input[4:])
			args := strings.Split(rawArgs, " ")

			for _, arg := range args {
				if slices.Contains(SHELL_COMMANDS, arg) {
					fmt.Printf("%s is a shell builtin\n", arg)
					break
				}
				if exe, err := commands.FindExecutable(os.Getenv("PATH"), arg); err == nil {
					fmt.Printf("%s is %s\n", arg, exe)
				} else {
					fmt.Printf("%s: not found\n", arg)
				}
			}

			continue
		}

		input = strings.TrimSpace(input)
		args := strings.Split(input, " ")
		// check if command is an executable to execute it
		exe, findErr := commands.FindExecutable(os.Getenv("PATH"), args[0])
		if findErr == nil {
			cmd := exec.Command(exe, args[1:]...)
			cmd.Args[0] = args[0]

			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			runErr := cmd.Run()
			if runErr != nil {
				fmt.Fprintln(os.Stderr, runErr)
			}

			continue
		}

		fmt.Printf("%s: command not found\n", input)
	}
}
