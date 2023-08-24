package main

import (
	"log"
	"net"
)

func main() {
	log.Println("Execution started")

	startListener("9090")
}

func startListener(port string) {
	listener, err := net.Listen("tcp", "0.0.0.0:"+port)

	if err != nil {
		log.Fatal("Error starting listener: ", err.Error())
	} else {

		channel, e := listener.Accept()
		defer channel.Close()

		if e != nil {
			log.Println("Error on a new channel: ", e.Error())
		}

		log.Println("New connection: ", channel.RemoteAddr().String())

	}
}
