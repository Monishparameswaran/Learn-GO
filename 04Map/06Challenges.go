package main
import "fmt"

/*
   Phonebook Application: Implement a simple phonebook application 
   where users can add, update, and search for contacts. Use a map to store the contact information.
*/

var phonebook=make(map[string]int);
func add(name string,num int){
	phonebook[name]=num;
}
func input(){
	var name string;
	var num int;
	fmt.Print("Enter name: ");
	fmt.Scanf("%v",&name);
	fmt.Println();
	fmt.Print("Enter number: ");
	fmt.Scanf("%v",&num);
	add(name,num)
	// fmt.Println(name, "->",num);
}

func call(name string){
	_,ok:=phonebook[name];
	if ok{
		fmt.Printf("calling %v --- %v \n",name,phonebook[name]);
	}else{
		fmt.Println("number not exist in phonebook");
	}
}
func main(){
	fmt.Println("welcome user !");
	input();
	var name string;
	fmt.Print("Enter name to call: ");
	fmt.Scanf("%v",&name);
	call("Monish");
}