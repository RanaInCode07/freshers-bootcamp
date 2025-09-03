package main

import(
	"fmt"
	"maps"
)

func main(){

	//map
	mp := make(map[string]int)
	mp["k1"] = 7
	mp["k2"] = 13

	fmt.Println("map:", mp)
	
	v1 := mp["k1"]
	v2:= mp["k3"] // non existing key returns zero value of value type
	fmt.Println("v1: ", v1)
	fmt.Println("v2: ", v2)

	fmt.Println("len:", len(mp))

	delete(mp,"k2")
	fmt.Println("map:", mp)

	mp["k3"] = 42
	mp["k2"] = 13

	clear(mp) // clearing a map
	fmt.Println("map:", mp)

	_, present := mp["k2"] // check if key is present in map
	fmt.Println("present:", present)

	mp["k1"] = 78
	var i int
	i, present = mp["k1"]
	fmt.Println("present:", present, i)

	n:= map[string]int{"foo":1, "bar":2} // map declared and initialized in one line
	fmt.Println("map:", n)

	 n2 := map[string]int{"foo": 1, "bar": 2}
     if maps.Equal(n, n2) { // comparing two maps
		 fmt.Println("n and n2 are equal")
	 }

    // range

	num := []int{2,3,4}
	sum:=0
	for _, n:= range num{
		sum += n
	}
	fmt.Println("sum:", sum)

	mp2:= map[string]int{"k1":1, "k2":2}
	for k,v:= range mp2{
		fmt.Println(k,v)
	}
	for k:= range mp2{ // iterating over keys only
		fmt.Println("key:", k)
	}

	for i,v:= range "ABCo"{ // iterating over string
		fmt.Println(i,v) // prints index and rune (unicode code point) value
	}


}