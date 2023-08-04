package main

import (
  "fmt"
  "os"
)

func errHandle(name error){
	if(name!=nil){
	fmt.Println("error status : Error occured !")
	fmt.Println("HINT: file might not be there");
	fmt.Println("the error says: ",name);}
}
/* By using this generalised code , we can avoid typing the errror handling code for each time such as 
if err!=nil{
		fmt.Printf("cannot open the file ! %s",err);
		fmt.Println();
	}
*/
func main(){
    _,err:=os.Open("user.txt")
	errHandle(err);
	_,err=os.Open("main.go"); // this will be successful so no error pops up
	errHandle(err);
}