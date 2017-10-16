package main

import (
	"fmt"
	"time"
)

func main(){
	jobs := make(chan int,100)
	tags := make(chan int,100)

	for i:=0;i<3 ;i++  {
		go work(i,jobs,tags)
	}

	for i:=0;i<5 ;i++  {
		jobs <- (i)
	}

	for i:=0;i<5 ;i++  {
		value := <- tags
		fmt.Println("tag finished ",value)
	}
}


func work(index int, jobs <-chan int, tags chan<- int){
	for i := range jobs{
		fmt.Println("worker",index,"start")
		time.Sleep(time.Second)
		fmt.Println("worker",index,"finish")
		tags <- (i)
	}
}



