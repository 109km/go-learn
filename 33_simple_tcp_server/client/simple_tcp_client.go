package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func errorHandler(err error) {
	fmt.Println("Error listening:", err.Error())
}

func main() {
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		errorHandler(err)
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	for {
		fmt.Println("What to send to the server? Type Q to quit.")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says: " + trimmedInput))
	}
}
