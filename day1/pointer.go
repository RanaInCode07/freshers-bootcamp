package main

import (
	"fmt"
)

func zeroOval(ival int) {
	ival = 0
}

func zeroPtr(iptr *int){
	*iptr = 0
}
func main(){
    i:=1
	fmt.Println("initial:", i)
	zeroOval(i)
	fmt.Println("after zeroOval:", i)
	zeroPtr(&i)
	fmt.Println("after zeroPtr:", i)
}