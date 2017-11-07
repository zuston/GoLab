package main

import "fmt"

func main(){
	var val = 100
	swap(val)
	fmt.Println(val)
	swapPointer(&val)
	fmt.Println(val)

	person := Person{
		name : "zuston",
		age : 10,
	}

	person.age = 20
	fmt.Println(person.age)
	perss := &person
	perss.age = 100
	fmt.Println(perss.age)
}


func swap(val int){
	val += 1
}

func swapPointer(val *int){
	*val += 1
}


type Person struct{
	name string
	age int
}