package commands

import (
	"helpers"
	"os"
)

type Ls struct {
	Command string
}

func (instance Ls) Exec() (response string, err error) {
	separatedCommand := helpers.CommandsSplit(instance.Command)

	var files []os.DirEntry

	if len(separatedCommand) > 1 {
		files, _ = os.ReadDir(separatedCommand[1])
	} else {
		files, _ = os.ReadDir(".")
	}

	for _, file := range files {
		response += file.Name() + "\n"
	}

	return "\n" + response, err
}
