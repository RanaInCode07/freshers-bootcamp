package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
}

func newPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42 
	return &p // return address of person
}

func main(){
	fmt.Println(Person{"Bob", 20}) // has to be in order of declaration
	fmt.Println(Person{name:"Alice"}) // age will be zero value 0
	fmt.Println(Person{name:"Fred", age: 40}) // name and age can be in any order
	fmt.Println(Person{}) // all fields will be zero value
	fmt.Println(&Person{name: "Ann", age: 40}) // return pointer to struct
	fmt.Println(newPerson("Jon")) // call newPerson function which returns pointer to struct

	s := Person{name:"Sean", age:50}
	fmt.Println(s.name) // accessing field of struct using dot operator

	sp := &s // pointer to struct
	fmt.Println(sp.age) // accessing field of struct using dot operator
	                   // no need to dereference the pointer using * operator
	                   // as Go automatically dereferences the pointer
	                   // when we use dot operator to access field of struct
	sp.age = 51 // struct is mutable

	fmt.Println(sp, s)

	dog := struct{ // anonymous struct used when struct is used only once
		name string
		age int
	}{name: "doggo", age: 5} // struct literal]

	fmt.Println(dog)
}

