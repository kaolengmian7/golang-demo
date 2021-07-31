package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type TestSVC struct {
}

func (t *TestSVC) ADD(args, reply *int) error {
	*reply = args.A + args.B
	return nil
}

func main() {
	testSVC := new(TestSVC)
	err := rpc.Register(testSVC)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	err = http.Serve(listen, nil)
	if err != nil {
		fmt.Printf("%+v", err)
	}
}
