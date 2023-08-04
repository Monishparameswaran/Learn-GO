package main
import "fmt"

func getDay(number int){
	switch(number){
	case 0:
		fmt.Println("Sunday");
	case 1:
		fmt.Println("Monday");
	case 2:
		fmt.Println("Tuesday");
	case 3:
		fmt.Println("Wednesday");
	case 4:
		fmt.Println("Thursday");
	case 5:
		fmt.Println("Friday");
	case 6:
		fmt.Println("Saturday");
	default:
		if number>31{fmt.Println("Invalid day");break;}
		getDay(number-7);
	}
}
 
type Employeew struct{
      name string
}
type Student struct{
	id int
}
type Printer interface{
	print();
}
func (e Employee) print(){

}
func (s Student) print(){

}
func typed_switch(data interface{}){

	switch val:=data.(type){
	case int:
		fmt.Println("the given value is integer type",val);
	case string:
		fmt.Println("the given value is string",val);
	case float32,float64:
		fmt.Println("the given value is float",val);
	case Printer:
		fmt.Println("this is a valid type",val);
	default:
		fmt.Println("the given type is invalid");

	}
}
type Employee struct{}
func main(){
	getDay(38);
	getDay(12);
	getDay(0);
	var e Employee;

	typed_switch("monish")
	typed_switch(123);
	typed_switch(e)

	var monish=Employeew{name :"monish"};
	var dhakshin=Student{id:235};

	typed_switch(monish)
	typed_switch(dhakshin)
}