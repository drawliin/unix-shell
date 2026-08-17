package commands

import (
	"fmt"
	"strings"
	"github.com/chzyer/readline"
)

type BellCompleter struct {
	Base readline.AutoCompleter
}

var lastCompletionLine string

func (b BellCompleter) Do(line []rune, pos int) ([][]rune, int) {
	out, offset := b.Base.Do(line, pos)
	if len(out) == 0 {
		fmt.Print("\x07")
	}

	if len(out) > 1 {
		current := string(line[:pos])
		
		if lastCompletionLine != current {
			fmt.Print("\x07")
			lastCompletionLine = current
			return nil, 0
		}

		fmt.Print("\n")
    	fmt.Print(strings.Join(reconstructCmd(current, out), "  "))
    	fmt.Print("\n")
		fmt.Print("$ " + string(line))
		
		return nil, 0
	}

	return out, offset
}
