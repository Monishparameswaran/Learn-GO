package main
import "fmt"

// Try this get a value in one goroutine and send to goroutine called sqauare then the squared value should be printed using 
// printer go routine 

// the writing channel should always need to be closed
func main(){

	wire:=make(chan int)
	NumSquare:=make(chan int)
	done:=make(chan bool)
	go get(wire);   /// creating a go routine
	go Sqaure(wire,NumSquare);  
	go Printer(NumSquare,done);
	<-done
	fmt.Println("all done.!");
}

func get(input chan int ){
	for i:=1;i<=5;i++{
	   input<-i;
	}
	close(input)
}
func Sqaure(input chan int,NumSquare chan int){
	for val:=range input{
		NumSquare<-val*val;
	}
	close(NumSquare)
}
func Printer(output chan int,done chan bool){
	for value:=range output{
		fmt.Println(value);
	}
	done<-true
}