package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/Haolaoda1/GOLang-Lab/net/rpc/domain"
)

type TestSVC struct {
}

func (t *TestSVC) ADD(args *domain.Args, reply *int) error {
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
