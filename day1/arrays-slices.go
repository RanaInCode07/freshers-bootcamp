package main

import (
	"fmt"
	"slices"
)

func main(){

	// array
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 30
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	b := [5]int{1,2,3,4,5}
	fmt.Println("dcl:", b)
	c := [...]int{1,6,2,5,3} // size is inferred from number of initializers
	fmt.Println("dcl:", c)
    fmt.Println("len:", len(c)) // length of array
	var twoD [2][3]int
	for i:=0; i<2; i++{
		for j:=0; j<3; j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", twoD)

	// slices
	var s[] string // zero value is nil
	fmt.Println("emp:", s)
	s = append(s, "hello")
	s = append(s, "world")
	s = append(s, "!")
	s = append(s, "Welcome", "to", "Go") // can append multiple values at once
	fmt.Println("set:",s)
	
	var p []int = make([]int,5) // make function creates a slice of given length
	fmt.Println("dcl:", p)
	p[0] = 1
	p[1] = 2
	p[2] = 3
	fmt.Println("set:", p)

	l:= s[2:5] // slice from index 2 to 4
	fmt.Println(s)
	fmt.Println("slice:", l)
	m:= s[:5] // slice from start to index 4
	fmt.Println("slice:", m)
	n:= s[2:] // slice from index 2 to end
	fmt.Println("slice:", n)

	t1:= []string{"g","h","i"} // slice literal
	t2:= []string{"g","h","i"}
	if slices.Equal(t1, t2){ // comparing two slices
		fmt.Println("t1 and t2 are equal")
	}

	newtwoD := make([][]int, 3) // slice of slices
	for i := range 3{
        innerlen := i+1
		newtwoD[i] = make([]int, innerlen)
		for j:= range innerlen{
			newtwoD[i][j] = i+j
		}
	}
	fmt.Println("2d: ", newtwoD)
}
