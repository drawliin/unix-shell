package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, _ := reader.ReadString('\n');
		
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

		fmt.Printf("%s: command not found\n", strings.TrimSpace(command));
	}
}

func isExitCommand(s string) bool {
	if (strings.TrimSpace(s) == "exit") {
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