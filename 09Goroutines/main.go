package main

import (
	"fmt"
)

// func simple_delayprint(num int){
// 	time.Sleep(time.Second * 4);
// 	fmt.Println(num);
// }
// func simple_go_routine(){
// 	for i:=0;i<10;i++{
// 		go simple_delayprint(i);  // the function call becomes  go routine // you can see the program do not wait for go 
// 	 // routine execution since a main function itself a go routine it is independent and do not wait for another go routine
// 	}
	
// 	time.Sleep(10 * time.Second);
// }

func generator(num chan int){
	for i:=0;i<=10;i++{
		num<-i;
	}
	close(num);
}

func routines_using_channel(){
	number:=make(chan int)
	go generator(number);
	for val:=range number{
		fmt.Println("channel generated value",val);
	}
	
}
func main(){
	//simple_go_routine()

	routines_using_channel();
}