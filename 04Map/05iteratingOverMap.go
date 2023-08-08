package main
import "fmt"

func main(){
	

	var mymap=make(map[string]int);
	mymap["Monish"]=7
	mymap["Parameswara"]=13

    fmt.Println(mymap);

	for i,v:= range mymap{
		fmt.Println(i,"->",v);
	}

	}