package commands

import (
	"d3c/server/helpers"
	"log"
)

type Select struct {
	Command []string
}

func (instance Select) Exec() (response string, err error) {
	if len(instance.Command) > 1 {
		if helpers.AgentRegistration(instance.Command[1]) {
			helpers.SelectedAgent = instance.Command[1]
			response = "Agent selected: " + instance.Command[1]
		} else {
			response = "The selected agent is not in the field."
			log.Println("The selected agent is not in the field.")
			log.Println("To list agents in the field type: show agents")
		}
	} else {
		helpers.SelectedAgent = ""
		response = "No agent selected"
	}

	return response, err
}
