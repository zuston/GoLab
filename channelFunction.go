package main

import "fmt"

func main(){
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for  {
			i, more := <-jobs
			// 如果 go channel closed 结果则是 false
			if more {
				fmt.Println("receive job", i)
			}else {
				fmt.Println("all jobs")
				done <- true
				return
			}
		}
	}()

	for i:=0; i<4 ;i++  {
		jobs <- i
	}
	close(jobs)
	fmt.Println("send to jobs end!")
	<- done
}
