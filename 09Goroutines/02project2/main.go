package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(pipe chan int){

	for i:=0;i<=10;i++{
		time.Sleep(2*time.Second);
		fmt.Println("From producer ",i);
		randomNumber:=rand.Intn(100);
		pipe<-randomNumber;    //sending the value to the channel
	}
	close(pipe);   // close the channel once it is over otherwise iot causes error
}

func Consumer(){
	pipe:=make(chan int);
	go Producer(pipe)
	for val:= range pipe{     /// consider the pipe as array where it sends the value continuosly the consumer will waiting
		                     ///for the producer until he send nil value
		fmt.Println("I m consumer i have got :",val);
	}
}

// remember we can create a channel between goroutine and normal function and as well as between two go routines 
func main(){
	go Consumer();

	time.Sleep(30 * time.Second);
}