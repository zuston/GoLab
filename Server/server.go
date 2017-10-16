package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"io"
)

func main(){
	tcpPort := flag.String("port","9090","port")
	flag.Parse()

	connection, err := net.Listen("tcp",":"+*tcpPort)
	if err!= nil {
		fmt.Println("error")
		os.Exit(1)
	}
	fmt.Println("tcp is open , listening port : ", tcpPort)
	defer connection.Close()

	for   {
		conn, err := connection.Accept()
		if err!= nil {
			fmt.Println("error")
			os.Exit(1)
		}
		fmt.Println("connection is built, remoteAddr : %v, localAddr : %v", conn.RemoteAddr(), conn.LocalAddr())
		go handleReques(conn)
	}
}
func handleReques(conn net.Conn) {
	defer conn.Close()
	for  {
		io.Copy(conn, conn)
	}
	
}