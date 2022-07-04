package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/minpeter/MSA-with-go/rpc_http/contract"
)

const port = 1234

type HelloWorldHandler struct{} // RPC 서버에 등록하기 위해 임의의 타입으로 정의

func (h *HelloWorldHandler) HelloWorld(args *contract.Args, reply *contract.Reply) error {
	reply.Message = "Hello " + args.Name
	return nil
}

func main() {

	StartServer()
}

func StartServer() {

	helloWorld := &HelloWorldHandler{}
	rpc.Register(helloWorld) // HelloWorldHandler 타입의 인스턴스를 서버에 등록
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		// log.Fatal(fmt.Sprintf("Unable to listen on given port: %s", err))
		log.Fatalf("Unable to listen on given port: %s", err)
	}

	log.Printf("Server string on port %v\n", port)

	http.Serve(l, nil)

	defer l.Close()

	for {
		conn, _ := l.Accept()

		defer conn.Close()

		go rpc.ServeConn(conn)
	}
}
