package main

import "fmt"

func method1(){

	fmt.Println("using make function.!");
	another_slice:=make([]int,2); // using make

	for i:=0;i<5;i++{
	    var num int;
		fmt.Print("enter the number:");
		fmt.Scan(&num);
		another_slice=append(another_slice,num )
	}
	
	fmt.Println(another_slice);
}

func method2(){
	fmt.Println("using normal declaration method ...!");
	slce:=[]int{};  // simply if a array is defined without initial value..
	for i:=0;i<5;i++{
	    var num int;
		fmt.Print("enter the number:");
		fmt.Scan(&num);
	slce=append(slce,num )
	}
	
	fmt.Println(slce);
}
func main(){



	method1()
	method2();
	
}