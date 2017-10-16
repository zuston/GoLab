package main

import (
	"fmt"
	"strings"
)

func Index(value []string, target string) int{
	for index ,value := range value{
		if value==target {
			return index
		}
	}
	return -1

}


func Include(value []string, target string) bool{
	return Index(value, target) > 0
}


func Any(value []string, comFunc func(string) bool) bool{
	for _ ,value := range value{
		if comFunc(value) {
			return true
		}
	}
	return false
}


func main() {
	values := []string{"zus","ppp"}
	res := Any(values, func(i string) bool {
		if strings.HasPrefix(i,"pp") {
			return true;
		}
		return false
	})
	fmt.Println(res)
}