package main

import "fmt"

func modify_array(a [3]int){
	a[0]=112;     // this is to demonstrate that changing passing arguments will pass as a value not as reference"
}
func main(){
	arr:=[3]int{1,2,4};
	fmt.Println("value of arr before passing : ",arr);
	modify_array(arr) 
	fmt.Println("value of arr after passing : ",arr);

	  // the scope of the array value is only restricted to its scope not to the entire program
}