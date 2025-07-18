package commands

import (
	"helpers"
	"strconv"
	"strings"
)

type Sleep struct {
	Command string
}

func (instance Sleep) Exec() (response string, err error) {
	separatedCommand := helpers.CommandsSplit(instance.Command)

	if len(separatedCommand) > 1 {
		timeStr := strings.TrimSpace(separatedCommand[1])
		_, err = strconv.Atoi(timeStr)
		if err != nil {
			return "Invalid sleep time: " + err.Error(), err
		} else {
			response = "Sleep time updated: " + timeStr + " seconds"
		}
	} else {
		response = "Usage: sleep <seconds>"
	}

	return response, err
}
