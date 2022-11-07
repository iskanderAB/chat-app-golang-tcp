package main

import (
	// "bufio"
	"bufio"
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	openConnetion = make(map[net.Conn]bool)
	newConnection = make(chan net.Conn)
	deadConection = make(chan net.Conn)
)

func main() {
	lr, err := net.Listen("tcp", ":8080")
	logFatal(err)
	defer lr.Close()

	go func() {
		for {
			conn, err := lr.Accept()
			logFatal(err)
			openConnetion[conn] = true
			newConnection <- conn

		}
	}()
	for {
		select {
		case conn := <-newConnection:
			go broadcastMessage(conn)
		case conn := <-deadConection:
			for connected := range openConnetion {
				if connected == conn {
					delete(openConnetion, conn)
					break
				}
			}
		}
	}

}

func broadcastMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		//1  loop through all  connections
		//2  send the message to all users connected except the sender and have fun ♥

		for connected := range openConnetion {
			if connected != conn {
				fmt.Println(msg)
				connected.Write([]byte(msg))
			}
		}
	}
	fmt.Println("fuuuuuck")
	deadConection <- conn
}
