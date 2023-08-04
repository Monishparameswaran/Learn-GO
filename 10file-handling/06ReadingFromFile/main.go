package main
import "fmt"
import "os"

import "bufio"
func main(){
	
	file,err:=os.OpenFile("myfile.txt",os.O_RDONLY | os.O_APPEND,0777)
	defer file.Close();

	if err!=nil{
		fmt.Println("cannot open the file");
	}

	scanner:=bufio.NewScanner(file);

	for scanner.Scan(){
		line:=scanner.Text()
		fmt.Println(line);
	}
	if scanner.Err()!=nil{
		fmt.Println("error while reading",scanner.Err());
	}
}