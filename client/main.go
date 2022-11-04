package main

import (
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn ,err := net.Dial("tcp",":8080")
	logFatal(err)
	defer conn.Close()

	fmt.Println("put ur name for conversation :)  ")
	var username string
	fmt.Scanln(&username)

	wlcMsg := fmt.Sprintf("welcome %s , to the chat , say Hi to your friends ♥ ", username)
	fmt.Println(wlcMsg)
	bytes := []byte(wlcMsg)
	fmt.Println(bytes)
	conn.Write(bytes)
}
