package main

import (
  "fmt"
  "os"
)

func openFile(a string)(error){
	
	_,err:=os.Open(a)
	return fmt.Errorf("hello user, The error says: %s ",err);
}
func main(){
    err:=openFile("monish.txt");
	if err!=nil{
		fmt.Println("error status: error occured");
		fmt.Println(err);
	}else{
		fmt.Println("error status: Not Exist");
		fmt.Println("opened the file");
	}
}