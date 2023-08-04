package main
import ("fmt")


func main(){
	arr:=[2][2]int{
		 {10,1},
		 {2,7}}

	// fmt.Println(arr); it prints the entire arrray in oneline

	for index,value:=range arr{
		fmt.Println(index," ",value);
	}
}