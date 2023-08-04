package main

import (
  "fmt"
  "os"
)


func main(){
    _,err:=os.Open("home.txt");
	if err!=nil{
		fmt.Printf("cannot open the file ! %s",err);
		fmt.Println();
	}
	_,err=os.Open("hre.txt");
	if err!=nil{
		fmt.Printf("cannot open this file ! %s",err);
		fmt.Println();
	}
}