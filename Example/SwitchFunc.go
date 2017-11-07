package main

import "fmt"

func main(){

	switchFunc := func(i interface{}) {
		switch t := i.(type) {
		case bool :
			fmt.Println("bool type")
		case int:
			fmt.Println("int type")
		default:
			fmt.Println(t)
		}
	}
	switchFunc(false)
	switchFunc(10)
}



