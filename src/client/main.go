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
	if len(arguments) == 2 {
		fmt.Println("Please provide a server and port number!")
		return
	}
	server := arguments[1]
	port := arguments[2]
	connection, err := net.Dial("tcp", server+":"+port)
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
