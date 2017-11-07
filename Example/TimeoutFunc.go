package main

import (
	"time"
	"fmt"
)

func main(){
	c1 := make(chan string, 1)
	go func(){
		time.Sleep(time.Second*2)
		c1 <- "zzzz!"
	}()


	select {
	case <- c1:
		fmt.Println("ok")
	case <-time.After(time.Second):
		fmt.Println("timeout")
		// 非堵塞，不需要等待
	default:
		fmt.Println("non blocking queue")
	}


	// 非堵塞的话，select 需要采用 default 状态

}