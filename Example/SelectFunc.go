package main

import (
	"time"
	"fmt"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(time.Second)
		c1 <- "done1"
	}()

	go func() {
		time.Sleep(time.Second*2)
		c2 <- "done2"
	}()

	for i:=0;i<2 ;i++  {
		select {
		case msg1:=<-c1:
			fmt.Println(msg1)
		case msg2:=<-c2:
			fmt.Println(msg2)

		}
	}
}