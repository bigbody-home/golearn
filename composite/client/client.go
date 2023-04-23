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
	defer con.Close()
	MustCopy(os.Stdout, con)

}

func MustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}

}
