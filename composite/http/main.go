package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Start ...")
	for {
		con, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)

		}
		go Handle(con)
	}

}

func Handle(con net.Conn) {
	defer con.Close()
	for {
		_, err := io.WriteString(con, fmt.Sprintf("time is %v", time.Now().Format("15:04:05\n")))
		if err != nil {
			return
		}
		time.Sleep(time.Second * 1)
	}

}
