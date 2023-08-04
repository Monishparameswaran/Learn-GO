package main
import "fmt"

func main(){
	arr:=[5]int{}
	for i:=0;i<5;i++{
		fmt.Print("enter the number:");
		fmt.Scan(&arr[i]);
		fmt.Println();
	}
	fmt.Println(arr);

	// what is slice

	
}