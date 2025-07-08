package helpers

import "strings"

func CommandsSplit(fullCommand string) (separateCommand []string) {
	parts := strings.Split(strings.TrimSuffix(fullCommand, "\n"), " ")
	// Remove argumentos vazios no final
	for len(parts) > 1 && parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}
	return parts
}
