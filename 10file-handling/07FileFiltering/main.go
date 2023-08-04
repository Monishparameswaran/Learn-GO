package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main(){

	file,err:=os.OpenFile("file.txt",os.O_RDONLY,0777)
	defer file.Close()
	if err!=nil{
		fmt.Println("cannot open, ",err);
	}
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line:=scanner.Text();
		if(strings.Contains(line,"name")){     
			fmt.Println(line);
		}
	}
   if scanner.Err()!=nil{
	fmt.Println("error occured");
   }

}