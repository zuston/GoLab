package main

import "fmt"

func main(){
	init := initSeq()
	fmt.Println(init())
	fmt.Println(init())

	recursion(1)
}


func initSeq() func() int{
	i := 1
	return func() int {
		i += 1
		return i
	}
}


func recursion(step int){
	if step==100 {
		fmt.Println("end")
		return
		}
	recursion(step+1)
}
