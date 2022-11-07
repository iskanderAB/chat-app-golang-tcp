package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	logFatal(err)
	defer conn.Close()

	fmt.Println("put ur name for conversation :)  ")
	var username string
	fmt.Scanln(&username)
	username = strings.Trim(username," \r\n")

	wlcMsg := fmt.Sprintf("welcome %s , to the chat , say Hi to your friends ♥ \n", username)

	fmt.Println(wlcMsg)
	go read(conn)
	write(conn,username)
}

func write(conn net.Conn, username string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		if err != nil{
			break
		}
		msg = fmt.Sprintf("%s:  %s \n",username, strings.Trim(msg," \n\r"))
		conn.Write([]byte(msg))	}
}

func read(conn net.Conn) {
	for{
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			conn.Close()
			fmt.Println("connection closed. ")
			os.Exit(0)
		}
		fmt.Println(msg)
		fmt.Println("*******************************")
	}
}