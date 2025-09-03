package main

import(
	"fmt"
	"errors"
)

func f(arg int) (int, error){
	if arg == 42{
		return -1, errors.New("can't work with 42")
	}
	return arg+3, nil
}
// predefined errors also known as sentinel errors
var errorOutOfTea = fmt.Errorf("out of tea") 
var errPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error{
	if arg == 2{
		return errorOutOfTea
	} else if arg == 4{
		return fmt.Errorf("making tea: %w", errPower) // wrapping error with %w verb
	}
	return nil
}

func main(){
    for _,i := range []int{7,42}{
		if result, err := f(i); err !=nil{
			fmt.Println("f failed:", err)
		} else{
			fmt.Println("f worked:", result)
		}
	}

	for i := range 5{
		if err := makeTea(i); err !=nil{
			if errors.Is(err, errorOutOfTea){ // checking for predefined error
				fmt.Println("We should buy more tea")
			} else if errors.Is(err, errPower){ // checking for wrapped error
				fmt.Println("We should check the power")
			} else{
				fmt.Println("unknown error")
			}
			continue
		}
		fmt.Println("here is your tea")
	}
}