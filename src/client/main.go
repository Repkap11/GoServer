package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Client")
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}
	address := ":" + arguments[1]
	connection, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	go listenForResponces(connection)

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		connection.Write([]byte(string(text)))
	}
}

func listenForResponces(connection net.Conn) {
	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		msg := strings.TrimSpace(string(netData))
		fmt.Println(msg)
	}

}
