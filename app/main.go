package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"shell/app/commands"

	"github.com/chzyer/readline"
)

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
}

func main() {
	// Build a completion tree for command names.
	completer := readline.NewPrefixCompleter(
		// Each PcItem is one command the user can complete with Tab.
		readline.PcItem("exit"),
		readline.PcItem("echo"),
		readline.PcItem("type"),
		readline.PcItem("pwd"),
		readline.PcItem("cd"),
	)

	// Create the readline instance and configure the shell prompt + autocomplete.
	reader, err := readline.NewEx(&readline.Config{
		Prompt:       "$ ",
		AutoComplete: commands.BellCompleter{Base: completer},
	})
	if err != nil {
		panic(err)
	}

	defer reader.Close()

	for {
		// Read one full input line from the interactive terminal.
		input, _ := reader.Readline()

		input = strings.TrimSpace(input)

		if len(input) == 0 {
			continue
		}

		cmdName, cmdArgs, stdRedirect := commands.SplitTokens(input)
		if cmdName == "" {
			continue
		}

		switch cmdName {
		// check exit command to exit the program
		case "exit":
			return

		// check echo command
		case "echo":
			commands.WriteOutput(stdRedirect, strings.Join(cmdArgs, " ")+"\n")

		// check pwd command
		case "pwd":
			pwd, _ := os.Getwd()
			commands.WriteOutput(stdRedirect, pwd+"\n")

		// check type command
		case "type":
			var output string

			for _, arg := range cmdArgs {
				if builtins[arg] {
					output = fmt.Sprintf("%s is a shell builtin\n", arg)
					continue
				}

				exe, err := exec.LookPath(arg)
				if err == nil {
					output = fmt.Sprintf("%s is %s\n", arg, exe)
				} else {
					output = fmt.Sprintf("%s: not found\n", arg)
				}
			}
			commands.WriteOutput(stdRedirect, output)

		// command to change directory
		case "cd":
			if len(cmdArgs) == 0 {
				continue
			}
			if len(cmdArgs) > 1 {
				fmt.Println("cd: too many arguments")
				continue
			}

			if err := os.Chdir(cmdArgs[0]); err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", cmdArgs[0])
			}

		// check if command is an executable to execute it
		default:
			cmd := exec.Command(cmdName, cmdArgs...)
			if err := commands.RunCommand(cmd, stdRedirect); err != nil {
				var exitErr *exec.ExitError
				if errors.As(err, &exitErr) {
					// The program was found and executed, but it exited with a failure code.
					// Do nothing here, because the program already wrote its own error to stderr.
					continue
				}
				fmt.Printf("%s: command not found\n", cmdName)
			}
		}
	}
}
