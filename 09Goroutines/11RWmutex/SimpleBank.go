package main

import (
	"errors"
	"fmt"
	"sync"
)

// this is a bank account project that shows how to withdraw ,checkbalance ,deposit money in your bank account
type Account struct{
	money int;
	sync.RWMutex

}

func (obj *Account) checkbalance()int{
	obj.RLock()
	defer obj.RUnlock()

	return obj.money
}

func (obj *Account) deposit(amount int)error{
	
	obj.Lock();  // write lock
	obj.money+=amount;
	fmt.Println("Money deposited",amount,"successfully..!")
	obj.Unlock(); // read lock
	if(obj==nil){
		return errors.New("could not find the bank account")
	}
	return nil;
}

func (obj *Account) withdraw(amount int){
	if obj.money>=amount{
		obj.Lock();
		obj.money-=amount;
		obj.Unlock();
	}
	fmt.Println("amount:",amount,"rs taken successfully ! balance: ",obj.checkbalance());
}
func main(){

	monish:=Account{};
	monish.deposit(100);
	fmt.Println(monish.checkbalance());
	monish.deposit(300);
	fmt.Println(monish.checkbalance());
	monish.withdraw(20)
	
}