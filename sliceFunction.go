package main

import "fmt"

func main(){
	a := []int{1,2,3,4}
	b := []int{5,6,7,8}

	x := append(a, b...)
	fmt.Println(x)

	fmt.Print(a)
}

type IntType int

func (a IntType) Write(c int) (n int , err error){
	return n+1, nil
}
