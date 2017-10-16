package main

import (
	"sort"
	"fmt"
)




type ByLength [] string

func (a ByLength)  Len() int{
	return len(a)
}

func (a ByLength) Swap(i,j int)  {
	a[i],a[j] = a[j],a[i]
}
func (a ByLength) Less(i,j int) bool{
	return len(a[i]) < len(a[j])
}

func main(){
	strs := []string{"csd","bsaaaaa","as"}
	sort.Sort(ByLength(strs))
	fmt.Println(strs)

	ints := []int{4,1,5}
	sort.Ints(ints)
	fmt.Println(ints)

}


