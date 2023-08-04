package main

import (
	"fmt"
)
func main(){

	var num int;
	fmt.Scan(&num);
	if(num==0){
		panic("zero error")
	}

	fmt.Println("i m executed after panic");
}