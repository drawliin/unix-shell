package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"shell/app/commands"
)

var builtins = map[string]bool{
	"exit": true,
	"echo": true,
	"type": true,
	"pwd":  true,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if len(input) == 0 {
			continue
		}

		cmdName, cmdArgs := commands.ArgsParser(input)
		
		switch cmdName {
		// check exit command to exit the program
		case "exit":
			return

		// check echo command
		case "echo":
			fmt.Println(strings.Join(cmdArgs, " "))

		// check pwd command
		case "pwd":
			pwd, _ := os.Getwd()
			fmt.Println(pwd)

		// check type command
		case "type":
			for _, arg := range cmdArgs {
				if builtins[arg] {
					fmt.Printf("%s is a shell builtin\n", arg)
					break
				}
				if exe, err := commands.FindExecutable(os.Getenv("PATH"), arg); err == nil {
					fmt.Printf("%s is %s\n", arg, exe)
				} else {
					fmt.Printf("%s: not found\n", arg)
				}
			}
		
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
			exe, findErr := commands.FindExecutable(os.Getenv("PATH"), cmdName)
			if findErr == nil {
				cmd := exec.Command(exe, cmdArgs...)
				cmd.Args[0] = cmdName

				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				runErr := cmd.Run()
				if runErr != nil {
					fmt.Fprintln(os.Stderr, runErr)
				}
			} else {
				fmt.Printf("%s: command not found\n", input)
			}
		}
	}
}
