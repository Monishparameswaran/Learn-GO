package main

import ("fmt"
      "time")

/*
Write a program that creates two goroutines. Each goroutine should
print a series of numbers from 1 to 5, with a slight delay between each number. 
Use the go keyword to launch the goroutines and observe the concurrent execution
*/

func Machine1(){
	for i:=1;i<=5;i++{
		time.Sleep(3 * time.Second);
		fmt.Println("I m Machine 1 printing: ",i);
	}
	fmt.Println("Machine 1 completed sir...");
}

func Machine2(){
	for i:=1;i<=5;i++{
		time.Sleep(2* time.Second);
		fmt.Println("From Machine 2 printing: ",i);
	}
	fmt.Println("Machine 2 completed sir...");
}
func main(){


	go Machine1(); // creating a go routine
	go Machine2(); // creating a go routine


	time.Sleep(30 * time.Second);  // try without this line you will get no output ... since main function executes faster than other function
}