package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	enter   = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)

func main() {
	con, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	go broadcast()
	for {
		lis, err := con.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleConn(lis)

	}
}
func broadcast() {
	//record all clients in connection client
	clients := make(map[client]bool)
	for {
		select {
		// if a client connect to server, then add client to clients
		case v := <-enter:
			clients[v] = true
		// if a client disconnect server, then delete client from clients
		case v := <-leaving:
			delete(clients, v)
			close(leaving)
		// if received message from message chan, then loop all client in clients and send message to client
		case msg := <-message:
			for c := range clients {
				c <- msg
			}
		}

	}
}
func handleConn(con net.Conn) {
	ch := make(chan string)
	go clientWriter(con, ch)
	who := con.RemoteAddr().String()
	ch <- "You are" + who
	message <- who + "has arrived"
	enter <- ch
	input := bufio.NewScanner(con) // block listen client input until client closed
	for input.Scan() {
		message <- who + ":" + input.Text()
	}
	leaving <- ch
	message <- who + "has left"
	con.Close()
}
func clientWriter(con net.Conn, msg chan string) {
	for v := range msg {
		fmt.Fprintln(con, v)
	}
}
