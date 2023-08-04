package main

import "fmt"


func modify_arr(a *[3]int){
	a[0]=122   /// here we are passing as a reference ,that is passing the original array itself
	a[2]=9898
}
func main(){
    arr:=[3]int{1,2,4};
	fmt.Println("value of arr before passing : ",arr);
	modify_arr(&arr) 
	fmt.Println("value of arr after passing : ",arr);  // you can see the change in the actual value of array
}