package main

import (
	"fmt"
	"os"
)

// func divide(x,y int) (int,error){
// 	if(y==0){
// 		return 0,fmt.Errorf("zero division error")
// 	}else{
// 		return x/y,nil
// 	}
// }
// func main(){
// 	val,Err:=divide(10,12);
// 	fmt.Println(val,Err);

// }

// func open(name string) error{
// 	_,err:=os.Open(name);

// 	if err!=nil{
// 		return fmt.Errorf("cannot open the file the error states ///:%s ////",err);
// 	}
// 	return errors.New("no error")}

func open() (s string) {

	
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println(p)
			s = "cannot open the file"
		}
	}()

	_, err := os.Open("file1.txt")
    
	if err != nil {
		panic("error cannot open the file")
	}
	fmt.Println("i m after panic") // cannot be executed since after panic

	return s

}

// this function returns a value ,but why ..?
/* lets understand its flow panic -> defer -> recover() -> from that the program starts recovering from the error... thinking
	  we have solved the error, then all the remaining code in the defer will be carried out and finally reaches the function return statement
	  ****note**** - the statements that are not in the defer function will not be executed even after the recover() process

	 what about the statements that are in main function that exist after function call ?
	    eg , func main(){
	        fmt.Println(open());
	        fmt.Println();
	        fmt.Println("i m after function call in main function");
}

     here all the statements after function call will be executed...!the rule what i have said for is only applicable to the current function scope where
	 the panic function is ?



 // how a function can return a value after panic ?
 // what recover actually does here ?
    it recovers our program from terminating by managing the error.


*/
func main() {
	fmt.Println(open())
	fmt.Println()
	fmt.Println("i m after function call in main function")
}
