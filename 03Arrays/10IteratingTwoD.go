package main
import "fmt"

func main(){
	arr:=[5][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9}}
	
	
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i]," ");
	}
   fmt.Println();
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr[1]);j++{
			fmt.Print(arr[i][j]," ");
		}
		fmt.Println();
		
	}
}