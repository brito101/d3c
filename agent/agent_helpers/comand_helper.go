package agent_helpers

func CommandValidation(command string) (id int) {
	mapping := map[string]int{
		"cd":     1,
		"ls":     1,
		"ps":     3,
		"pwd":    4,
		"whoami": 5,
	}

	id = mapping[command]

	return
}
