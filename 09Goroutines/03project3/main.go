package main

import (
	"fmt"
	"math/rand"
	"time"
)
var count int=2;

type Customers struct{
	name string;
	movie string;
	id int; 
}

func NewCustomers(n ,m string) *Customers{
	return &Customers{name:n,movie:m};
}

func Producer(id chan Customers){

	  count=count-1;
	   if count>=1{
		time.Sleep(time.Second *2)
		randomID:=rand.Intn(26372);
		cust:=Customers{id: randomID};
		id<-cust;
	   }else{
		cust:=Customers{id:-122};
		id<-cust;
	   }
	   close(id);
		
}

func Print(obj *Customers){
	if obj.id>=0 {
		fmt.Println("----------------------------------------");
	fmt.Println("Your ticket was booked successfully..!");
	fmt.Println("Name: ",obj.name);
	fmt.Println("Movie: ",obj.movie);
	fmt.Println("ID: ",obj.id);
	fmt.Println("---------------------------------------------");
	}else{
		fmt.Println("----------------------------------------");
	fmt.Println("Failed");
	fmt.Println("Name: ",obj.name);
	fmt.Println("Movie: ",obj.movie);
	fmt.Println("ID: ",obj.id);
	fmt.Println("---------------------------------------------");
	}
}
func Consumer(obj *Customers){

	id:=make(chan Customers);
	go Producer(id);
	obj.id=(<-id).id;
	go Print(obj);
}
func main(){

	names:=[]*Customers{};
     var temp_name string;
	 var temp_movie string;
	for i:=0;i<=2;i++{
		fmt.Println("Welcome user,Enter your name");
		fmt.Scan(&temp_name)
		fmt.Println("Enter the movie name");
		fmt.Scan(&temp_movie)
		names=append(names,NewCustomers(temp_name,temp_movie))
		go Consumer(names[i]);
	}
	time.Sleep(time.Second *30)
}