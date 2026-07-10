package commands

import (
	"fmt"
	"os"
)

func WriteOutput(std stdPath, output string) {
	if std.path == "" {
		fmt.Print(output)
		return
	}

	file, err := os.Create(std.path)
	if err != nil {
		fmt.Printf("%s: %v\n", std.path, err)
		return
	}

	if std.stdout {
		if _, err = file.WriteString(output); err != nil {
			fmt.Printf("%s: %v\n", std.path, err)
		}
		return
	}

	fmt.Print(output)
}
