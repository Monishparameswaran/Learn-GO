package main

import (
	"fmt"
	"os"
)

func errorHandle(err error) {
	if err != nil {
		fmt.Println("an error occured.! ", err)
	}
}
func main() {
	file, err := os.Open("myfile.txt") // open a file
	defer file.Close()
	errorHandle(err)
}
