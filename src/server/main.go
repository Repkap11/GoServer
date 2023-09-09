package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type ClientList struct {
	lock    sync.RWMutex
	clients []net.Conn
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	address := ":" + arguments[1]
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	var clientList ClientList

	for {
		connection, err := listener.Accept()
		clientList.clients = append(clientList.clients, connection)
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(connection, &clientList)
	}
}

func handleConnection(connection net.Conn, clientList *ClientList) {
	defer connection.Close()
	fmt.Printf("Serving %s\n", connection.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		msg := strings.TrimSpace(string(netData))
		if msg == "STOP" {
			break
		}
		fmt.Println(msg)
		result := "Someone said:" + msg + "\n"
		clientList.lock.RLock()
		for _, client := range clientList.clients {
			client.Write([]byte(string(result)))
		}
		clientList.lock.RUnlock()
	}
}
