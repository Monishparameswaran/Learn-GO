package main
import "fmt"

type Employee struct{
	name string;
	emid int;
}

func setEmployee(value string,id int) *Employee{
	return &Employee{name:value,emid:id};
}
func main(){

	
	musk:=setEmployee("Elon musk",8978300);
	fmt.Println(musk);
	// var tata struct;  will not work
    // tata=setEmployee("Ratan tata",97340783);  
	
}