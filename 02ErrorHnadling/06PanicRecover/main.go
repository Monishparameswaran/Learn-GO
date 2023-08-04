package main

import (
	"fmt"
)

func checkvalidName(name string) {
	if name != "Monish" {
		panic("it is not a valid name")
	}
}
func main() {

	var name string
	fmt.Print("please enter name:")
	fmt.Scanf("%s", &name)

	
	
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("an error occured")
		}
	}()  ///putting this defer function after the function call will not work
	checkvalidName(name)
	fmt.Println("i can go head")
}
