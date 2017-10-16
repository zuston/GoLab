package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"bufio"
	"math/rand"
)

var connContainer map[string] net.Conn

func main(){
	
	connContainer = make(map[string]net.Conn)

	host := flag.String("host","","such as localhost , 198.235,63.0")
	port := flag.String("port","9090","connection port")

	flag.Parse()

	openAddr := *host + ":" + *port
	serverConnection, err := net.Listen("tcp",openAddr)
	if err!=nil {
		fmt.Println("server connection error")
		os.Exit(1)
	}
	defer func() {
		serverConnection.Close()
		fmt.Println("server connection closed")
	}()

	for{
		conn, err := serverConnection.Accept()
		if err!=nil {
			continue
		}
		key := conn.RemoteAddr().String()+string(rand.Int())
		connContainer[key] = conn
		go handle(conn,key)
	}
}
func handle(conn net.Conn, key string) {
	//fromAddr := conn.RemoteAddr().String()
	fromAddr := key
	defer func() {
		conn.Close()
		fmt.Println("the connection closed , ip :", fromAddr)
	}()

	buffer := bufio.NewReader(conn)
	for {
		msg, err := buffer.ReadString('\n')
		if err != nil {
			fmt.Println("connection err, ip : ", fromAddr)
			break
		}
		fmt.Println("msg is ", msg)
		broadcast(msg, fromAddr)
	}
}
func broadcast(msgStr string, fromAddrId string) {
	msg := []byte(string(fromAddrId+msgStr))
	for key, value := range connContainer{
		if key==fromAddrId {
			continue
		}
		value.Write(msg)
	}
}

