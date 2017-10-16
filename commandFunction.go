package main

import (
	"flag"
	"fmt"
)

func main(){
	var param1 *string = flag.String("localPort","200","local port")
	var param2 *string = flag.String("remotePort","300","remote port")

	flag.Parse()

	fmt.Println(*param1)
	fmt.Println(*param2)
}
