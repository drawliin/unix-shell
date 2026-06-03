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
		if isExitCommand(command) {
			break
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