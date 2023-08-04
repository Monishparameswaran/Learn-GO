package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
func tell(num int) string{
	switch(num){
	case 2:
		return "maths"
	case 3:
		return "science"
	case 4:
		return "cpp"
    case 6:
		return "java"
	case 5:
		return "os"
	}
	return "nil"

}
func main(){


	file,err:=os.OpenFile("Ass1.txt",os.O_RDONLY,0777);
	if err!=nil{
		fmt.Println("cannot open the file",err);
	}

	scanner:=bufio.NewScanner(file);

	var name string;
	fmt.Print("enter the name :");
	fmt.Scan(&name);
	var total int=0;

	for scanner.Scan(){
		line:=scanner.Text();
		if strings.Contains(line,name){
			fmt.Println();
			fmt.Println("Student name: ",name);
			field:=strings.Fields(line)
			for i:=2;i<=6;i++{
				data:=fmt.Sprintf("%s		 %s",tell(i),field[i]);
				fmt.Println(data);
				num,err:=strconv.Atoi(field[i])
				if err==nil{
					total+=num;
				}
			}
			fmt.Println();
			fmt.Println("total mark:	",total);
			fmt.Println();
			}
			
		}
	}