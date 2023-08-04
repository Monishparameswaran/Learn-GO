package main
import "fmt"

func main(){
	num:=[10]int{1,2,3};
	num[0]=10;
	fmt.Println(num); // Prints the entire array
	fmt.Println();
	for i:=0;i<len(num);i++{
		fmt.Print(num[i]," ");
	}
	fmt.Println();
}