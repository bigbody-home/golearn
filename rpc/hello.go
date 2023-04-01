package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (h *HelloService) Hello(request string, response *string) error {
	*response = "Hello" + request
	return nil
}

func (h *HelloService) Hello2(request string, response *string) error {
	*response = "你好啊" + request
	return nil
}
func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		con, err := listen.Accept()
		log.Println(con.RemoteAddr())
		if err != nil {
			log.Fatalln(err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(con))
	}

}
