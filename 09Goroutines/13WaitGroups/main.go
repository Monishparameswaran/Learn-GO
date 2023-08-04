package main

import "fmt"
import "sync"



var num int=0;
var wg sync.WaitGroup
func increment(){
	num++;
	wg.Done();
}

func main(){
	for i:=0;i<5;i++{
		go increment();
		wg.Add(1);
	}
	wg.Wait();
	fmt.Println(num);
}