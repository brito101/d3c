package main

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"global"
	"log"
	"net"
	"os"
	"time"
)

var (
	message   = global.Message{}
	heartBeat = 30
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

		gob.NewEncoder(channel).Encode(message)

		gob.NewDecoder(channel).Decode(message)

		time.Sleep(time.Duration(heartBeat) * time.Second)
	}

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
