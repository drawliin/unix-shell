package commands

import (
	"fmt"
	"os"
)

func WriteOutput(path string, output string) {
	if path == "" {
		fmt.Print(output)
		return
	}

	if err := os.WriteFile(path, []byte(output), 0644); err != nil {
		fmt.Printf("%s: %v\n", path, err)
	}
}
