package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	con, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, con) //from server receiving information to client stdout
		log.Print("done")
		done <- struct{}{}

	}()
	mustCopy(con, os.Stdout) //from client sending some information to server
	con.Close()
	<-done
}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
