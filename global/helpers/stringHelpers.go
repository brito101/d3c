package helpers

import "strings"

func CommandsSplit(fullCommand string) (separateCommand []string) {
	separateCommand = strings.Split(strings.TrimSuffix(fullCommand, "\n"), " ")
	return separateCommand
}
