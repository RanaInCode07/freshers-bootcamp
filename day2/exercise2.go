package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

    const totalStudents = 200
	var wg sync.WaitGroup
	var mu sync.Mutex

	var totalRatings int
	var totalResponses int


func main(){
	fmt.Println("Class monitor is collecting the rating of students")

	for i:=0; i<200; i++{
		wg.Add(1) // added a go routine so wait group is added by 1
		go GetRating(i)
	}
	wg.Wait() // waiting for all go routine to finish before running the main go routinge
	fmt.Println("All ratings have been collected") 
	avgRating := float64(totalRatings)/ float64(totalResponses) // calculating average ratings
	fmt.Println("Average rating are:", avgRating)
}

func GetRating(studentId int){
	    defer wg.Done() //reducing the counter of wait group by 1 after executing this go routine 
        randomdelay := time.Duration(rand.Intn(500)) * time.Millisecond //calculating randome delay
        time.Sleep(randomdelay) // sleeping for that random delay
		ratings := rand.Intn(5) + 1 // calculating ratings using rand
		mu.Lock() // using mutex to lock in the use of shared resources and variable to avoid race condition
		totalRatings += ratings // updating shared variables
		totalResponses++
		mu.Unlock()  //unlocking the mutex to unlock the shared resources that are being used
		fmt.Printf("Student %d provided a rating of %d\n", studentId, ratings)
}

