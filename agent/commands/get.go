package commands

import (
	"helpers"
	"os"
)

type Get struct {
	Command string
}

func (instance Get) Exec() (response string, err error) {
	separatedCommand := helpers.CommandsSplit(instance.Command)

	if len(separatedCommand) > 1 && len(separatedCommand[1]) > 0 {
		// Check if file exists and is readable
		_, err := os.ReadFile(separatedCommand[1])
		if err != nil {
			return "File read error: " + err.Error(), err
		}

		// File validation successful - actual file reading is handled by agent
		response = "File validation successful: " + separatedCommand[1]
	} else {
		response = "Usage: get <file>"
	}

	return response, nil
}
