# Switch

```
package main
import "fmt"

	switch(number){
	case 0:
         statements;
         // break is there by default
	case 1:
		statements
	case 2:
		statements 
	default:
        // it is like a else part executed at last after checking all case
	}
}
 
func main(){
	getDay(38);
	getDay(12);
	getDay(0);
}

```

type switches
In Go, a typed switch, also known as a type switch, is a variation of the switch statement that allows you to check the dynamic type of an interface value and perform different actions based on its underlying type. Here's how typed switches work:



```

switch x := value.(type) {
case Type1:
    // code to execute when value has Type1
case Type2,Type3:
    // code to execute when value has Type2
default:
    // code to execute when value has none of the specified types
}

```
switch can take cases with interface as a type.

```
package main
import "fmt"


type Employee struct{
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
        fmt.
	default:
		fmt.Println("the given type is invalid");

	}
}

func main(){
	

	typed_switch("monish")
	typed_switch(123);
	typed_switch(e)

	var monish=Employeew{name :"monish"};
	var dhakshin=Student{id:235};

	typed_switch(monish)
	typed_switch(dhakshin)
}
