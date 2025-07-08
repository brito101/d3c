package agent_helpers

func CommandValidation(command string) (id int) {
	mapping := map[string]int{
		"cd":     1,
		"ls":     2,
		"ps":     3,
		"pwd":    4,
		"whoami": 5,
	}

	id, ok := mapping[command]
	if !ok {
		return -1
	}
	return id
}
