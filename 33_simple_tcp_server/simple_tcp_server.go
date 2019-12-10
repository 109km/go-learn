package main

import (
	"fmt"
	"net"
)

func errorHandler(err error) {
	fmt.Println("Error listening:", err.Error())
}

func main() {
	fmt.Println("Starting the server ...")

	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		errorHandler(err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			errorHandler(err)
			return
		}
		fmt.Printf("%v\n", string(buf[:len]))

	}
}
