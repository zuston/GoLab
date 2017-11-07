package main

import (
	"net/rpc"
	"os"
	"fmt"
)

import "util"

func main(){
	client, err := rpc.DialHTTP("tcp","192.168.0.1:9999")
	if err!=nil {
		os.Exit(1)
	}
	args := &util.Args{1,2}
	var reply int
	e := client.Call("Arith.Multiply",args,&reply)
	if e!= nil {
		os.Exit(1)

	}
	fmt.Print("result : ",reply)
}
