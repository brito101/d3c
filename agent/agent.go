package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Execution started")
	log.Println("ID: ", generateID())
}

func generateID() string {
	hostname, _ := os.Hostname()
	time := time.Now().String()

	hash := md5.New()

	hash.Write([]byte(hostname + time))

	return hex.EncodeToString(hash.Sum(nil))
}
