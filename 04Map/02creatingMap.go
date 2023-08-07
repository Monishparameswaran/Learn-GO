package main
import "fmt"

func main(){
	store:=map[int]string{
		1:"Giri",
		2:"Jagan",
		3:"Monish"};
	fmt.Println(store);

	fmt.Println(store[3]); // Monish
}