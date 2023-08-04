package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("enter the file name to create a file")

	var name string
	fmt.Scan(&name)

	/*
	   Creating at a particular location is easy if know the address of that location ,in below example we have used
	   sprintf which creates a new address string according to the filename given 
	   
	*/
	destination:=fmt.Sprintf("/home/monish/golang-projects/learn-go/10file-handling/%s.txt",name);
	file, err := os.Create(destination) //create a file using os.Create

	if err != nil {
		fmt.Println("ERROR , cannot create a file.", err)
	} else {
		fmt.Println("file created successfully..!")
	}

	defer file.Close()
}
