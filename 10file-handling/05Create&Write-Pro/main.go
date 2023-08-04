package main

import (
	"bufio"
	"fmt"
	"os"
)
func main(){
	file,err:=os.OpenFile("myfile21.txt",os.O_WRONLY|os.O_APPEND|os.O_CREATE,777);
	if err!=nil{
		fmt.Println("an error occured ",err);
	}


	obj:=bufio.NewWriter(file);

	var data string="hello monish"
	_,err=obj.WriteString(data);

	if err!=nil{
		fmt.Println("an error occured...!");
	}else{
		fmt.Println("written data successfully");
	}

	// since here buffer is used to make sure that all the data is written onto the file, flushing the buffer transfer all the data completely
	err=obj.Flush()
	if err!=nil{
		fmt.Println("error occured while flushing the file");
	}
	err=file.Close()
	if err!=nil{
		fmt.Println("error while closing the file");
	}
}