package main

import ("fmt")

func main(){
	/* 
	What is slice 
	*/
	name:=[]string{"hello"};
	name=append(name,"Monish");
	name=append(name,"Jagan");
	name=append(name,"Giri");
	fmt.Println(name);
}