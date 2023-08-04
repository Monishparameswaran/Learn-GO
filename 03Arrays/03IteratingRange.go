package main
import "fmt"

func main(){
	var names=[10]string{"SteveJobs","Elon","Billgates","Tata"}

	for i,v :=range names{
		fmt.Println(i," ",v);
	}
}