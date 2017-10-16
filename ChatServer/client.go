package main

import (
	"flag"
	"net"
	"fmt"
	"os"
	"bufio"
)

var errorSemp chan bool

func main(){

	errorSemp = make(chan bool)

	host := flag.String("host","","such as localhost , 198.235,63.0")
	port := flag.String("port","9090","connection port")

	flag.Parse()

	openAddr := *host + ":" + *port
	clientConnection, err := net.Dial("tcp",openAddr)
	if err!=nil {
		fmt.Println("clientConnection  error")
		os.Exit(1)
	}
	defer func() {
		clientConnection.Close()
		fmt.Println("clientConnection closed")
	}()


	go receiveMsg(clientConnection)
	for{
		var msg string
		//fmt.Print(clientConnection.LocalAddr().String())
		fmt.Scanln(&msg)
		if msg=="quit" {
			errorSemp <- true
			break
		}
		val := []byte(msg+"\n")
		clientConnection.Write(val)
	}

	<-errorSemp
}
func receiveMsg(conn net.Conn) {
	buffer := bufio.NewReader(conn)
	for{
		msg,err := buffer.ReadString('\n')
		if err!=nil {
			errorSemp <- true
			break
		}
		fmt.Println(msg)
	}
}
