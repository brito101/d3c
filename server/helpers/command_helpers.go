package helpers

import "global"

func CommandValidation(command string) (id int) {
	mapping := map[string]int{
		"show":   1,
		"select": 2,
		"send":   3,
		"get":    4,
	}

	id, ok := mapping[command]
	if !ok {
		return -1
	}
	return id
}

func AddCommandToAgent(command global.Command, agentID string) {
	for i, v := range FieldAgents {
		if v.AgentID == agentID {
			FieldAgents[i].Commands = append(FieldAgents[i].Commands, command)
		}
	}
}
