package commands

import (
	"fmt"

	"github.com/chzyer/readline"
)

type BellCompleter struct {
	Base readline.AutoCompleter
}

func (b BellCompleter) Do(line []rune, pos int) ([][]rune, int) {
	out, offset := b.Base.Do(line, pos)
	if len(out) == 0 {
		fmt.Print("\x07")
	}
	return out, offset
}
