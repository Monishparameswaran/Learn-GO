package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var users []*Account;
type Account struct {
	name  string
	money int
	sync.Mutex
}

func (obj *Account) checkBalance() int {
	return obj.money
}

func (obj *Account) deposit(amount int) {
	time.Sleep(time.Second * 3)
	obj.Lock()
	defer obj.Unlock()
	obj.money += amount
	fmt.Println("Money deposited successfully, current balance:", obj.checkBalance())
}

func (obj *Account) serveMoney(pipe chan int, amount int) {
	obj.Lock()
	defer obj.Unlock()
	obj.money = obj.money - amount
	pipe <- amount
	close(pipe)
}

func (obj *Account) transferMoney(amount int, user *Account) {
	pipe := make(chan int)
	obj.Lock()
	defer obj.Unlock()

	if obj.money > amount {
		go obj.serveMoney(pipe, amount)
		val := <-pipe
		user.deposit(val)
		fmt.Println("Money transfer successful")
	} else {
		fmt.Println("Insufficient funds!")
	}
}

func (obj *Account) withDraw(amount int) {
	obj.Lock()
	defer obj.Unlock()

	if obj.money > amount {
		obj.money -= amount
		fmt.Println("Withdrawal successful, debited amount:", amount)
	} else {
		fmt.Println("Insufficient funds!")
	}
}

func (obj *Account) message(num int, amount int, msg string) {
	time.Sleep(time.Second * 3)
	fmt.Println()
	fmt.Println("---------------------------------------------------------------------")
	fmt.Println("Name:", obj.name)

	switch num {
	case 2:
		fmt.Println("Request: Deposit")
		fmt.Println("Money deposited:", amount)
		fmt.Println("Balance:", obj.money)
		fmt.Println("Status:", msg)
	case 3:
		fmt.Println("Request: Withdraw")
		fmt.Println("Balance:", obj.money)
		fmt.Println("Debited amount:", amount)
		fmt.Println("Status:", msg)
	case 4:
		fmt.Println("Request: Transfer")
		fmt.Println("Money transferred:", amount)
		fmt.Println("Balance:", obj.money)
		fmt.Println("Status:", msg)
	case 5:
		fmt.Println("Account created successfully")
	default:
		fmt.Println("-------------------------------------------------")
	}

	fmt.Println()
	fmt.Println()
}

func checkDataBase(name string) int {
	time.Sleep(time.Second * 2)
	for index, usr := range users {
		if usr.name == name {
			return index
		}
	}
	return -1
}

func newConstructor(user string) *Account {
	return &Account{name: user}
}

func createNewAccount(name string ,wg *sync.WaitGroup) {
	defer wg.Done()
	account := newConstructor(name)
	users = append(users, account)
	account.message(5, 0, "Account created successfully")
}

func route(num int,temp int, name string, wg *sync.WaitGroup) {
	time.Sleep(time.Second *10)
	defer wg.Done()
	var money int
	var index int
	var user string

	switch num {
	case 1:
		users[index].checkBalance()
	case 2:
		fmt.Println("Enter money:")
		fmt.Scan(&money)
		users[index].deposit(money)
	case 3:
		fmt.Println("Enter money:")
		fmt.Scan(&money)
		users[index].withDraw(money)
	case 4:
		fmt.Println("Enter user account name to transfer money:")
		fmt.Scan(&user)
		index = -1
		for val, usr := range users {
			if usr.name == user {
				index = val
			}
		}
		if index == -1 {
			fmt.Println("user does not exist")
		} else {
			users[temp].transferMoney(money, users[index])
		}

	case 5:
		createNewAccount(name,wg)
	}

}

func main() {
	
	wg.Add(5)

	for i := 0; i < 5; i++ {
		var name string
		var num int

		fmt.Println("Enter username:")
		fmt.Scan(&name)

		fmt.Println("Checking...")
		var index int = checkDataBase(name)
		if index != -1 {
			fmt.Println("Enter 1 to check balance")
			fmt.Println("Enter 2 to deposit")
			fmt.Println("Enter 3 to withdraw")
			fmt.Println("Enter 4 to transfer money")
			fmt.Println("Enter 5 to create a new account")
			fmt.Scan(&num)
			fmt.Println("Please wait, we are processing your request...")
			fmt.Println()
		
			go route(num,index, name, &wg)
		} else {
			fmt.Println("User does not exist. Creating a new account...")
			createNewAccount(name, &wg)
		}
	}

	wg.Wait()
}
