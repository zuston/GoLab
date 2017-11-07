package main

import "fmt"

type Student struct {
	name string
	age int
}

type Employee struct {
	pname string
	page int
}

func (a Student) Say(){
	fmt.Println("hello student")
}
func (a Student) Sing(){
	fmt.Println("sing student")
}


func (a Employee) Say(){
	fmt.Println("hello employee")
}

func (a Employee) Sing(){
	fmt.Println("sing employee")
}

type Men interface {
	Sing()
	Say()
}

func Meansure(g Men){
	g.Say()
	g.Sing()
}

func main()  {
	stu := Student{
		name : "zston",
		age : 100,
	}
	Meansure(stu)
}
