package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	for i :=0; i<5; i++ {
		fmt.Println("Value of i is", i)
	} // these braces are mandatory in go even for single statement body
    sum := 1
	for ; sum<10; { // the init and post statements are optional
		            // omitting the loop condition creates an infinite loop
		sum += sum
	}

	v := 9
	if v < 10 {
		fmt.Println("v is less than 10")
	}
	if a:=10; a<20 { // we can also have initialization statement in if in go
		             // variables initialized in if are also available in else block
		fmt.Println("a is less than 20")
	}
	fmt.Println("Sum is", sum)


	// switch block

	switch os:= runtime.GOOS; os{ // like if switch can also have initialization statement
		                          // runtime.GOOS gives the OS name and os is a variable
	case "linux":
		fmt.Println("Linux OS")
	case "darwin":
		fmt.Println("Mac OS")
	default:
		fmt.Println(os)
	}

	today := time.Now().Weekday()


	switch time.Tuesday{
	case today + 0: fmt.Println("Today")
	case today + 1: fmt.Println("Tomorrow")
	case today + 2: fmt.Println("In two days")
	default: fmt.Println("Too far away")
	}
    t := time.Now()
	switch { // switch without an expression is like switch true
		     // used to write long if-else-if chains
	case t.Hour() < 12:
		fmt.Println("Good morning")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")}

		//defer statement
        // all arguments to deferred function are evaluated immediately
		// but the function call is not executed until the surrounding function returns
		// deffered function calls are executed in LIFO order pushed into stack while deferring
		defer fmt.Println("World")
		fmt.Print("Hello ")
}