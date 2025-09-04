package main

import (
	"fmt"
	"sync"
)

// func main(){

// 	mp := make(map[string] int)
// 	v:= [] string{"quick", "brown", "fox", "lazy", "dog"}
// 	for _, i:= range v{
//           for j:=0; j<len(i); j++{
// 			 mp[string(i[j])]++
// 		  }
// 	}
// 	fmt.Println(mp)
// }

func CountFrequency(s string, resultChan chan<- map[rune]int, wg *sync.WaitGroup){
     defer wg.Done() // decrement the wait group counter when the function returns 

	 localmap := make(map[rune] int)

	 for _,r := range s{
		localmap[r]++
	 }
   resultChan <- localmap // send the local map to the channel
}

func main(){
	inputStrings := [] string{"quick", "brown", "fox", "lazy", "dog"}
// use wait group to wait for all the goroutine to finish
	var wg sync.WaitGroup
	//creating a channel to recieve frequency map from goroutines
	resultChan := make(chan map[rune]int, len(inputStrings))
// to count each string character in different go routing
	for _,s:= range inputStrings{
		wg.Add(1)
		go CountFrequency(s, resultChan, &wg)
	}
// after all go routine are finished close the channel
	go func(){
       wg.Wait()
	   close(resultChan)
	}()

	finalmap:= make(map[string] int)
	for localmap:= range resultChan{
		for char, count:= range localmap{
			finalmap[string(char)] += count
		}
	}
	fmt.Println(finalmap)


}