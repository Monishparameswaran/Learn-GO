package main
import "fmt"

type Student struct{
	name string;
	regno int;
}

/*
 following is the constructor which is nothing but a function which is used to
  assign value for struct elements.
*/

func setStudent(name string,no int) Student{
	return Student{name:name,regno: no};
}

func main(){
	var monish Student;

	monish.name="Monish"
	monish.regno=69937

    var jagan Student;    // declaring variable jagan
	jagan=setStudent("Jagan",647392)  ///assigning values using constructor

	
	giri:=setStudent("Giri",970936) 

	fmt.Println(giri);
	fmt.Println(monish);
	fmt.Println(jagan);

}