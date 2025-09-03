package main

import(
	"fmt"
	"math"
)


type Vertex struct{
    X,Y float64
}

func (v Vertex) Abs() float64{ // abs is a method on Vertex type with receiver v
	                           // Abs can only be called with Vertex type
	return math.Sqrt(v.X*v.X +v.Y*v.Y)
}

//using pointer receiver

func(v *Vertex) Scale(f float64){
   v.X = v.X*f
   v.Y = v.Y*f
}

func(v *Vertex) AbsPointer() float64{
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
func main(){
	v:=Vertex{3,4}
	fmt.Println(v.Abs()) // calling method Abs on Vertex type
    a := &Vertex{3,4} // a is pointer to Vertex type
	fmt.Printf("Before Scaling: %+v, Abs: %v\n", a, a.AbsPointer())
	a.Scale(5) // calling method Scale on Vertex type
	fmt.Printf("After Scaling: %+v, Abs: %v\n", a, a.AbsPointer())
}

//You can only declare a method with a receiver whose type is defined in the same package as the method.
//You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
//functions with a pointer argument must take a pointer:
//while methods with pointer receivers take either a value or a pointer as the receiver when they are called:
//That is, as a convenience, Go interprets the statement v.Scale(5) as (&v).Scale(5) since the Scale method has a pointer receiver
