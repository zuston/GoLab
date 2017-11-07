package main

import (
	"fmt"
)

func main(){
	var a = [5]int{1,2,3,4,5}
	fmt.Println(a)
	var b = make([]string, 5)
	fmt.Println(b)
	var c [5] int
	fmt.Println(c)

	fmt.Println(len(b))
	b[0] = "hello"
	b[4] = "sd"
	fmt.Println(b)
	b = append(b, "zuston")
	fmt.Println(b)

	l := b[1:]
	fmt.Println(l)


	mapper := map[string]int{"ann":1993,"zuston":1994}
	fmt.Println(len(mapper))

	for key,value := range mapper{
		fmt.Println(key,value)
	}

	// 判断一个mapper的key是否存在
	v, prs := mapper["ann"]
	if prs {
		fmt.Println(v)
	}else{
		fmt.Println("the key not existed")
	}

	// range
	var mmapper = make(map[string]int)
	mmapper["p"] = 1
	mmapper["q"] = 2
	mmapper["e"] = 4
	for k,v := range mmapper{
		fmt.Printf("%s:%d\n",k,v)
	}


	for k,v := range "gooo"{
		fmt.Println(k,v)
	}

	_, dd := plus(1,2)
	fmt.Print(dd)

	sum(1,2,3,4)
	sum(1)
}


func plus(a int, b int) (int, int){
	return a+b, a-b
}



func sum(nums ... int){
	summer := 0
	for _, value := range nums {
		summer += value
	}
	fmt.Println("sum = ",summer)
}