package agent_helpers

const (
	CMD_CD     = 1
	CMD_LS     = 2
	CMD_PS     = 3
	CMD_PWD    = 4
	CMD_WHOAMI = 5
	CMD_SEND   = 6
	CMD_GET    = 7
	CMD_SLEEP  = 8
)

func CommandValidation(command string) (id int) {
	mapping := map[string]int{
		"cd":     CMD_CD,
		"ls":     CMD_LS,
		"ps":     CMD_PS,
		"pwd":    CMD_PWD,
		"whoami": CMD_WHOAMI,
		"send":   CMD_SEND,
		"get":    CMD_GET,
		"sleep":  CMD_SLEEP,
	}

	id, ok := mapping[command]
	if !ok {
		return -1
	}
	return id
}
