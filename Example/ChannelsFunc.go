package main

import (
	"fmt"
	"time"
)

func main(){
	msg := make(chan string)

	go func() {
		time.Sleep(10000)
		msg <- "ping"
	}()

	mssg := <- msg
	fmt.Println(mssg)

	mmm := make(chan string, 2)
	mmm <- "sd"
	mmm <- "ddd"
	fmt.Println(<-mmm)


	done := make(chan bool, 1)
	go workder(done)

	<- done
}



func workder (done chan bool){
	fmt.Println("working ....")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
// 控制 channel 流向方向
func ping(pings chan <- string, msg string){
	
}