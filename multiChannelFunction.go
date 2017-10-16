package main

import (
	"fmt"
	"time"
)

func main(){
	semphore := make(chan int,1)

	end := make(chan bool,3)

	for i:=0;i<3 ;i++  {
		go func() {
			v,more := <- semphore
			if more {
				fmt.Println("runtime number ",i,v)
				end <- true
				return
			}
			end <- true
			fmt.Println("have no more")
		}()
	}

	time.Sleep(time.Second*3)
	semphore <- 1

	close(semphore)

	for i:=0;i<3 ;i++  {
		<- end
	}
}
