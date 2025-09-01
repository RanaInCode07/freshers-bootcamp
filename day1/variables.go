package main

import(
    "fmt"
)

var count, name int // package level variable declaration

func main(){
	var j int // local variable declaration function level
    var i, k = 1, 2 // variable declaration with initialization
	                // type is inferred from the value assigned
	a := 3 // used in place of var with implicit type
	       // can only be used inside functions
	       // cannot be used for package level variable declaration
	// j := 4 // this will throw error as j is already declared in this scope
 
    fmt.Println(j, i , k, a)
}