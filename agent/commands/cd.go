package commands

import (
	"helpers"
	"os"
)

type Cd struct {
	Command string
}

func (instance Cd) Exec() (response string, err error) {

	response = "Current directory change success!"
	separatedCommand := helpers.CommandsSplit(instance.Command)

	if len(separatedCommand[1]) > 0 {
		err = os.Chdir(separatedCommand[1])

		if err != nil {
			response = "Directory change error: " + err.Error()
		}
	}

	return response, err
}
