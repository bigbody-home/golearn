package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var host = flag.String("host", "", "host info")
	var port = flag.String("port", "", "port info")
	var user = flag.String("user", "", "user info")
	flag.Parse()
	if *host == "" || *port == "" || *user == "" {
		log.Fatalln("参数不能为空")
		os.Exit(1)
	}

	fmt.Printf("USER=%s,HOST=%s,port=%s\n", *user, *host, *port)
}
