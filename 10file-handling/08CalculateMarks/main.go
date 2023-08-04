package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func main(){

	file,err:=os.OpenFile("file.txt",os.O_RDONLY,0777)
	if err!=nil{
		fmt.Println("cannot open the file");
	}

	scanner:=bufio.NewScanner(file);
	var total int;
	for scanner.Scan(){
		line:=scanner.Text();
		field:=strings.Fields(line);

		if len(field)>=3{
			num,err:=strconv.Atoi(field[2]);

			if err==nil{
				total+=num;
			}
		}
	}
	fmt.Println("the total mark:",total);

}