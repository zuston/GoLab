package main

import (
	"fmt"
)

func main(){
	val, err := g(1)
	fmt.Println(val)
	if err!=nil {
		fmt.Println(err.(*Err).msg)
	}
}


func g(val int) (int, error){
	if val==1 {
		return val, &Err{val:val,msg:"the val is false"}
	}
	return val, nil
}


type Err struct {
	val int
	msg string
}

func (err *Err) Error() string{
	return fmt.Sprintf("%s-%s",err.val, err.msg)
}