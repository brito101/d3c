package commands

import (
	"helpers"
	"os"
)

type Cd struct {
	Command string
}

func (instance Cd) Exec() (response string, err error) {
	separatedCommand := helpers.CommandsSplit(instance.Command)

	if len(separatedCommand) > 1 && len(separatedCommand[1]) > 0 {
		err = os.Chdir(separatedCommand[1])
		if err != nil {
			return "Directory change error: " + err.Error(), err
		}
		return "Current directory change success!", nil
	} else {
		return "Usage: cd <directory>", nil
	}
}
