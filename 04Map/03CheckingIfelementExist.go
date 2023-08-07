package main
import ("fmt")

func main(){
	var mymap=make(map[string]string)
	mymap["name"]="Monish";
	mymap["interest"]="open source and devops"
	_,ok:=mymap["name"] // _ used in as variable to store the values when we are not going to use the value anymore,golang will not ask us to use this variable
	
	// _,check:=mymap["interest"]

	_,mon:=mymap["info"] 

	// here ok,mon are just varibles that hold boolean values 

	if ok{
		fmt.Println("value exist");
	}else{
		fmt.Println("value not exist for name field");
	}

	if mon{
		fmt.Println("the value exist for info");
	}else{
		fmt.Println("value not exist for key info");
	}

}