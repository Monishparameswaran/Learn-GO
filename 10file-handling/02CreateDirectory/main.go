package main
import "fmt"
import "os"

func main(){

	fmt.Print("enter name of the directory to create: ");
	var name string;
	fmt.Scan(&name);
	err:=os.Mkdir(name,0757);
	if err!=nil{
		fmt.Println("cannot create a directory..! ",err);
	}
	fmt.Println("created succesfully");
}