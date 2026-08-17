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

	flags := os.O_CREATE | os.O_WRONLY

	if std.appendEOF {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	file, err := os.OpenFile(std.path, flags, 0644)
	if err != nil {
		fmt.Printf("%s: %v\n", std.path, err)
		return
	}
	
	defer file.Close()

	if std.stdout {
		if _, err = file.WriteString(output); err != nil {
			fmt.Printf("%s: %v\n", std.path, err)
		}
		return
	}

	fmt.Print(output)
}

func reconstructCmd(current string, arr [][]rune) []string {
	var out []string
	for _, word := range arr {
		out = append(out, current + string(word))
	}
	return out
}