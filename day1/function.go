package main

import(
	"fmt"
)

func add( a int, b int) int{ // type is written after the variable name
	return a+b
}

func mul(a, b int) int{ // if consecutive variables are of same type
	                     // we can write type only once
	return a*b
}

func swap(x, y string) (string, string){ // function can return multiple values
	return y, x
}

func split(sum int) (x, y int) { // named return values
	x = sum * 4 / 9
	y = sum - x
	return // naked return, return statement without any variables
	       // only used in functions with named return values
}

// variadic functions
// can take variable number of arguments
func variadic(nums ...int) int {
	total := 0
	for _,num:= range nums{ // range returns both index and value
		                     // we can use blank identifier _ to ignore one of them
		total += num
	}
	return total
}

//closures
func intSeq() func() int { // function returns another function
	                       // returned function forms a closure
	                       // as it references the variable i
	                       // defined in the enclosing function
						   // closures used for data hiding
	i:=0
	return func() int{
		i++
		return i
	}
}

func fact(n int) int {
	if n==0{
		return 1
	}
	return n*fact(n-1)
}

func main(){
	sum := add(23, 87)
	fmt.Println("Sum is", sum)
	product := mul(12, 6)
	fmt.Println("Product is", product)
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(23))
	fmt.Println(variadic(5,9,2,8,3))

	nextInt := intSeq() // nextInt is a function
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(fact(5))

	var fib func(n int) int // function variable declaration
	fib = func(n int) int{
		if n<2{
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}