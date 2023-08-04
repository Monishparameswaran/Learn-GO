package main

import "fmt"

func main(){
	name:=make([]string,0,10)

	name=append(name, "m")
	name=append(name, "o")
	name=append(name, "n")
	name=append(name, "i")

	fmt.Println(name[1:3]);
	

}