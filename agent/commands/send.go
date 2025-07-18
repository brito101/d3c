package commands

import (
	"helpers"
)

type Send struct {
	Command string
}

func (instance Send) Exec() (response string, err error) {
	separatedCommand := helpers.CommandsSplit(instance.Command)

	if len(separatedCommand) > 1 && len(separatedCommand[1]) > 0 {
		// This command is handled by the server's file handling logic
		// which receives the file from the agent
		response = "File upload command received: " + separatedCommand[1]
	} else {
		response = "Usage: send <file>"
	}

	return response, err
}
