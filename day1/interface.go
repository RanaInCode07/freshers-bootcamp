package main

import(
	"fmt"
	"math"
)

type Abser interface{ // defining interface
	Abs() float64
}


type MyFloat float64 // user defined type

type Vertex1 struct{ // user defined struct type
	X,Y float64
}

func (f MyFloat) Abs() float64{ // implementing Abs method with receiver of type MyFloat of Abser interface
      if f<0{
		return float64(-f)
	  }
	  return float64(f)
}



func (v Vertex1) Abs() float64{ // implementing Abs method with receiver of type Vertex of Abser interface
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main(){
	var a Abser
	f:=MyFloat(-math.Sqrt2)
	v:= Vertex1{3,4}
    
	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
	a=v // a Vertex implements Abser
	fmt.Println(a.Abs())
}

//nil interface 

// An interface value that holds neither value nor concrete type is called a nil interface.
// The nil interface value is represented internally by a pair of nil pointers.
// A nil interface value can be assigned to any interface type, but it does not implement any methods, so calling a method on a nil interface causes a run-time panic.	

// empty interface

// An interface that specifies zero methods is known as the empty interface.
// The empty interface is written as interface{}.
// Since every type implements at least zero methods, all types implement the empty interface.
// Values of type interface{} can hold values of any type.
// Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
