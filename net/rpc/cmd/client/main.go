package main

import (
	"fmt"
	"net/rpc"

	"github.com/Haolaoda1/GOLang-Lab/net/rpc/domain"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Printf("%+v", err)
	}
	args := new(domain.Args)
	args.B = 1
	args.A = 1
	var reply int
	//同步调用
	err = client.Call("TestSVC.ADD", args, &reply)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Println(reply)
}
