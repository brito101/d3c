package global

type Message struct {
	AgentID       string
	AgentHostname string
	AgentCWD      string
	Commands      []Command
}
