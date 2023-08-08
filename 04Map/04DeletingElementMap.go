package main
import "fmt"

func main(){
			
	var mymap=make(map[string]int);
	mymap["Monish"]=12
	mymap["jagan"]=13
    fmt.Println("before deletion",mymap);

	delete(mymap,"Monish");
	
	fmt.Println("after deletion",mymap);
}