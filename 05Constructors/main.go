package main
import (
	"fmt"
)

type student struct{
	name string
	num int
}

func NewConstructor(name string,n int)*student{
	return &student{name:name,num:n}
}
func main(){

	monish:=NewConstructor("monish",46478);
	fmt.Println(monish)
}