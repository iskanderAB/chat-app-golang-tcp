package main

import (
	"fmt"
	"log"
	"net"
)

func logFatal(err *error) {
	if *err != nil {
		log.Fatal(err)
	}
}


var (
	openConnetion = make(map[net.Conn]bool)
	newConnection = make(chan net.Conn)
	deadConection = make(chan net.Conn)
)
// var newConnection
func main() {
	lnr, err := net.Listen("tcp",":8080")
	logFatal(&err)

	defer lnr.Close()

	go func(){
		for {
			conn, err := lnr.Accept()
			logFatal(&err)
			openConnetion[conn]= true
			newConnection <- conn
		}
	}()

	fmt.Println((<-newConnection).LocalAddr().String())
	
    fmt.Scanln() // bech sel server yo93ed yekhdem ! 
}