package main
import ("fmt")


func main(){
	arr:=[3][2]int{
		 {10,1},
		 {2,7},
		{3,4}}

	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr[0]);j++{
			fmt.Print(arr[i][j]," ");
		}
		fmt.Println();
	}
}
