package commands

import (
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/chzyer/readline"
)

func MakeCompleter() readline.AutoCompleter {
	items := []readline.PrefixCompleterInterface{
		readline.PcItem("exit"),
		readline.PcItem("echo"),
		readline.PcItem("type"),
		readline.PcItem("pwd"),
		readline.PcItem("cd"),
	}

	for _, name := range pathExecutables() {
		items = append(items, readline.PcItem(name))
	}

	return readline.NewPrefixCompleter(items...)
}

func pathExecutables() []string {
	seen := map[string]bool{}
	var names []string

	for _, dir := range filepath.SplitList(os.Getenv("PATH")) {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			info, err := entry.Info()
			if err != nil {
				continue
			}

			name := entry.Name()
			if isExecutable(name, info) {
				name = strings.TrimSuffix(name, filepath.Ext(name))
				if !seen[name] {
					seen[name] = true
					names = append(names, name)
				}
			}

		}
	}

	sort.Strings(names)
	return names
}

func isExecutable(name string, info os.FileInfo) bool {
	if runtime.GOOS != "windows" {
		// Unix executable bit.
		// 0111 is octal which maps to 001 001 001 and checks only the execution bit for all groups
		return info.Mode()&0111 != 0
	}

	ext := strings.ToLower(filepath.Ext(name))
	for _, pathext := range filepath.SplitList(os.Getenv("PATHEXT")) {
		if ext == strings.ToLower(pathext) {
			return true
		}
	}

	return false
}
