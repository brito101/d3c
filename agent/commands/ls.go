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

	if len(separatedCommand) > 1 && len(separatedCommand[1]) > 0 {
		files, err = os.ReadDir(separatedCommand[1])
		if err != nil {
			return "Directory not found: " + separatedCommand[1], err
		}
	} else {
		files, err = os.ReadDir(".")
		if err != nil {
			return "Error reading current directory", err
		}
	}

	for _, file := range files {
		response += file.Name() + "\n"
	}

	return "\n" + response, err
}
