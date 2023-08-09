package main
import "fmt"

type Student struct{
	name string;
	regno int;
}

func main(){
	var monish Student;  // creating object for student struct
	
	monish.name="Monish Parameswara"
	monish.regno=3487252;

	var giri Student;
	
	giri.name="Giri Dharan"
	giri.regno=78671921;

}