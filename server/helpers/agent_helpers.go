package helpers

import "global"

var (
	FieldAgents   = []global.Message{}
	SelectedAgent = ""
)

func AgentRegistration(agentID string) (registered bool) {
	registered = false

	for _, v := range FieldAgents {
		if v.AgentID == agentID {
			registered = true
		}
	}

	return registered
}

func CommandResponse(message global.Message) (contains bool) {
	contains = false

	for _, v := range message.Commands {
		if len(v.Response) > 0 {
			contains = true
		}
	}

	return contains
}

func AgentFieldPosition(agentID string) (position int) {
	for i, v := range FieldAgents {
		if v.AgentID == agentID {
			position = i
		}
	}

	return position
}
