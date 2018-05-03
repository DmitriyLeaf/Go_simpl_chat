package main

import (
	"fmt"
	"net"
	"bufio"
	"container/list"
	)

var clients *list.List

func Client(client net.Conn, ){
	bufReader := bufio.NewReader(client)

	for {
		rline, err := bufReader.ReadString('\n')
		
		if err != nil {
			fmt.Println("Cen not read!", err)
			client.Close()
			break
		}

		fmt.Println(client.RemoteAddr(), ": ", string(rline))

		for i := clients.Front(); i != nil; i = i.Next() {
			fmt.Fprint(i.Value.(net.Conn), client.RemoteAddr())
			fmt.Fprint(i.Value.(net.Conn), ": ")
			fmt.Fprint(i.Value.(net.Conn), rline)
		}
	}
}

func main() {
	fmt.Println("Run server:")

	clients = list.New()

	server, err := net.Listen("tcp", ":8080")

	if err != nil {
		fmt.Println("Error: %s", err.Error())
		return
	}

	for {
		client, err := server.Accept()

		if err != nil {
			fmt.Println("Error: %s", err.Error())
			client.Close()
			continue
		}

		clients.PushBack(client)
		fmt.Println("Connected: ", client.RemoteAddr())
		
		go Client(client)
	}

}