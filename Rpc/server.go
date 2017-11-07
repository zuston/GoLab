package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// rpc 透明调用方法
func main() {
	// 注册一个 struct
	arith := new(Arith)
	rpc.Register(arith)
	// http 通信
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":9999")
	if e != nil {
		log.Fatal("error : ", e)
	}

	go http.Serve(l, nil)
	for {
		fmt.Print("")
	}
}
