package main

import (
	"flag"
	"net"
	"fmt"
	"os"
)

func main(){
	var port *string = flag.String("port","9090","remote port")
	flag.Parse()
	conn, err := net.Dial("tcp",":"+*port)
	if err!=nil {
		fmt.Println("error")
		os.Exit(1)
	}

	defer conn.Close()
	tag := make(chan string)
	go handleWrite(conn, tag)
	go readAction(conn, tag)
	<-tag
	<-tag
}
func readAction(conn net.Conn, strings chan string) {
	buffer := make([]byte,102400)
	lenn, err := conn.Read(buffer)
	if err!=nil {
		fmt.Println("error read")
		os.Exit(1)
	}
	fmt.Println(string(buffer[:lenn-1]))
	strings <- "read done"
}
func handleWrite(conn net.Conn, strings chan string) {

	for i:=0;i<10 ;i++  {
		_, err := conn.Write([]byte("hello world "+string(i)))
		if err!=nil {
			fmt.Println("error write")
			break
		}
	}
	strings <- "write done"
}

