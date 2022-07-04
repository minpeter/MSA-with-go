package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/minpeter/MSA-with-go/rpc/contract"
)

const port = 1234

func main() {
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port)) // RPC 서버 접속
	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer client.Close()

	// 동기 호출
	args := &contract.Args{Name: "peter"}
	var reply contract.Reply

	err = client.Call("HelloWorldHandler.HelloWorld", args, &reply)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(reply.Message)

	//비동기 호출
	args.Name = "minpeter"
	sumCall := client.Go("HelloWorldHandler.HelloWorld", args, &reply, nil) //고루틴으로 호출
	<-sumCall.Done
	fmt.Println(reply.Message)
}
