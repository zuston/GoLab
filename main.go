package main

import (
	"fmt"
	"math/rand"
	"math"
	"runtime"
)


var name string
var age int

func main(){
	//defer fmt.Println("hello ")

	fmt.Print("hello world")
	fmt.Println()
	fmt.Println(rand.Intn(12))
	fmt.Println(math.Pi)
	name = "ziuston"
	change(name)
	fmt.Println(name)


	const PATH = "world hello"
	fmt.Println(PATH)

	loop()
	switchFunction()

	deferFunction()


	var p *int

	a := 1
	p = &a
	fmt.Println(*p)
	*p = 34
	fmt.Println(a)

	var cc[2]string
	cc[0] ="sd"

	fmt.Println(len(cc))

	// 数组传递引用
	pp := make([]int,4)
	pp[0] = 34
	testArr(pp)
	fmt.Println(pp)
}


func add(x, y int) int {
	return x+y
}


func swap(x, y string) (string, string){
	return y, x
}


func change(x string) string {
	return x+"sdsd"
}

func loop(){
	sum := 0
	for i:=0;i<10 ;i++  {
		sum += i
	}
	fmt.Println(sum)
}

func switchFunction(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func deferFunction(){
	for i:=0; i<10 ; i++  {
		defer fmt.Println(i)
	}
}


type Person struct {
	name string
	age int
}


func buildPerson() Person{
	return Person{
		name : "zuston",
		age: 23,
	}
}

// 构造 slice 数组
func sliceFunction(){
	a := make([]int, 6)
	a[0] = 2
}


func testArr(pp []int){
	pp[0] = 23
}

var personMap map[string]Person


