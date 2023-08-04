package main
import (
	"fmt"
	"time"
)

func main(){

	Value:=make(chan int)
	go give(Value);  
	go get(Value);


	// both go routines run concurrently since channel is used in between them to communicate the value ,what happens here is
	// consider two buses going in the road where the people from one bus is getting to another through the pipe on the go 
	time.Sleep(time.Second *30);
}
func give(out chan int){
	for i:=0;i<=10;i++{
		time.Sleep(time.Second *2);
		out<-i;
	}
	close(out)
}
func get(in chan int){
	for val:=range in{
		fmt.Println(val);
	}
}