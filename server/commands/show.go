package commands

import (
	"d3c/server/helpers"
	"log"
)

type Show struct {
	Command []string
}

func (instance Show) Exec() (response string, err error) {
	if len(instance.Command) > 1 {
		switch instance.Command[1] {
		case "agents":
			for _, v := range helpers.FieldAgents {
				response += "ID Agent: " + v.AgentID + " -> " + v.AgentHostname + "@" + v.AgentCWD + "\n"
			}
		default:
			response = "The selected parameter does not exist"
			log.Println("The selected parameter does not exist")
		}
	} else {
		response = "Usage: show <parameter>"
	}

	return response, err

}
