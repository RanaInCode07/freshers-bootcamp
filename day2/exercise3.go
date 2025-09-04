package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BankAccount struct{
	balance int
    mu      sync.Mutex
}

func (acct *BankAccount) Deposit(amount int, wg *sync.WaitGroup){
	defer wg.Done()

	acct.mu.Lock()
	defer acct.mu.Unlock()
	acct.balance += amount
	fmt.Printf("Amount Deposited: %d, Current Balance: %d\n", amount, acct.balance)
}

func (acct *BankAccount) Withdraw(amount int, wg *sync.WaitGroup){
	defer wg.Done()
	
	acct.mu.Lock()
	defer acct.mu.Unlock()
	if acct.balance >= amount {
	   acct.balance -= amount
	   fmt.Printf("Amount Withdrawn: %d, Current Balance: %d\n", amount, acct.balance)
	} else{
		fmt.Printf("Withdrawal of %d failed. Insufficient Balance, Current Balance: %d\n", amount, acct.balance)
	}	
}

func main(){

	rand.Seed(time.Now().UnixNano())

    account := BankAccount{balance:500}
	var wg sync.WaitGroup

	const numTransactions = 10

	wg.Add(numTransactions)

	for i:=0; i<numTransactions; i++{
        if i%2 == 0{
			go account.Deposit(rand.Intn(100)+1, &wg)
		} else{
			go account.Withdraw(rand.Intn(300)+1, &wg)
		}
	}
	wg.Wait()
	fmt.Println("All transaction finished\n")
	fmt.Printf("Final Balance: %d\n", account.balance)

}




// var mu1 sync.Mutex
// var wg1 sync.WaitGroup

// func deposit(balance *int, deposit int){
// 	 wg1.Add(1)
// 	 defer wg1.Done()
// 	 mu1.Lock()
//      *balance += deposit
// 	 fmt.Println("Balance credited, Remaning balance:", *balance)
// 	 mu1.Unlock()
// }

// func withdraw(balance *int, deposit int){
// 	 wg1.Add(1)
// 	 defer wg1.Done()
// 	 if *balance >= deposit{
// 		mu1.Lock()
//         *balance -= deposit
// 		fmt.Println("Balance debited, Remaninig balance:", *balance)
// 		mu1.Unlock()
// 	 } else {
// 		 fmt.Errorf("You don't have enough money to withdraw")
// 	 }
// }

// func main(){
//     balance := 500
// 	go deposit(&balance, rand.Intn(300))
//     go withdraw(&balance, rand.Intn(300))
// 	wg1.Wait()
// 	fmt.Println("Balance finished all thread executed")
// }