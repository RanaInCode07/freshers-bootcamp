package main

import (         // writing import in this way is called factored import
	"fmt"
	"math/rand" // by convention package name is same as last element of 
	           // import path in this case it is rand package
	"math"
)

func main(){
	fmt.Println("My favorite number is ", rand.Intn(10))
	fmt.Println(math.Pi). // accessing Pi constant, in go capitalized names are exported from a package
}