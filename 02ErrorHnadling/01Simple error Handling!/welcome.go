package main
import (
	"fmt"
	"os"
)

func main(){
    _,err:=os.Open("file1.txt");
   if err!=nil{
	fmt.Println("file is not there !");
   }
}