package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("myfile.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}

	fmt.Println("File deleted successfully.")
}
