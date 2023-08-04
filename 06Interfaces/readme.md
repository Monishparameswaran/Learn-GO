# Interface

An interface is a collection of method signatures that define a set of behaviors. It specifies what methods a type must have to be considered as implementing the interface. Interfaces provide a way to define abstractions and decouple code from specific implementations.

if muliple structures has same function then type interface can define that common function signature in it

``` 
package main

import (
	"fmt"
)

type Employee struct{
	salary int;
	name string;
}
type Student struct{
	name string;
}

func (e Employee) print(message string){
	fmt.Printf("%s to the employee :%s\n",message,e.name );
}

func (s Student) print(instruction string){
	fmt.Printf("%s welcome on board :%s\n",instruction,s.name);
}

type Printer interface{
	print(string);
}

func main(){

	var monish Printer;
	var musk Printer;
	monish=Student{name: "monish",}
	musk=Employee{name:"Elon musk",salary:1000000000000000000,}
	monish.print("must adhere the rules and regulations");
	musk.print("nothing comes from luck");
}


```

## How Interfaces are useful

 Interfaces in Go provide a powerful mechanism for abstraction, polymorphism, dependency injection, testing, and code modularity. They promote flexible and modular code design, making it easier to write maintainable and extensible Go programs.

 Abstraction and Decoupling: Interfaces in Go allow you to define abstractions and decouple code from specific implementations. By programming to interfaces rather than concrete types, you can write code that is more flexible and independent of the underlying implementations.

Polymorphism: Go interfaces enable polymorphism, allowing you to write code that can work with different types as long as they satisfy the interface requirements. This promotes code reuse and flexibility, as you can swap out implementations without affecting the code that uses the interface.

Dependency Injection: Interfaces play a crucial role in dependency injection scenarios. By defining interfaces for dependencies, you can inject different implementations at runtime, making your code more testable, modular, and flexible.

Mocking and Testing: Interfaces are essential for unit testing and mocking in Go. By defining interfaces for external dependencies, you can create mock implementations to isolate and test specific parts of your code. This facilitates easier testing and verification of interactions between components.

API Design: Interfaces are used to define clear and consistent APIs in Go libraries and frameworks. By exposing well-defined interfaces, library authors provide a contract for users to interact with their code. Users can then implement those interfaces to extend or customize the library's behavior.

Standard Library Usage: The Go standard library heavily utilizes interfaces. Many standard library functions and methods accept interfaces as arguments, allowing you to provide custom implementations. This enables greater flexibility and extensibility in using the standard library.

Code Modularity and Reusability: By programming to interfaces, you create code that is modular, reusable, and loosely coupled. Interfaces help in separating concerns, promoting cleaner code architecture and easier code maintenance.

Dynamic Behavior: Go interfaces, coupled with type assertions and type switches, enable dynamic behavior. You can check whether a value satisfies an interface and extract the underlying concrete type, allowing for runtime polymorphism and flexible behavior.

## how go interfaces are different from interfaces of other programming languages like java?

>In go interfaces are implicit where in java it is explicit

>Go has an empty interface interface{}, which can hold values of any type,Java does not have a direct equivalent of an empty interface.

>Go provides type assertions and type switches to check whether a value satisfies an interface and to extract the concrete type behind an interface. This allows for runtime type checks and dynamic behavior. Java, on the other hand, uses explicit casting to check and convert types, which is less flexible and requires explicit handling of potential exceptions.

## Interfaces can take any value of any type

This made interface more powerful 

```
package main
import "fmt"


func check(data interface{}){
	

	if val,ok:=data.(string);ok{
		fmt.Println("it is a string ,no error",val);
	}else{
		fmt.Printf("Unsupported type\n");
	}
}
func main(){
	check(1234); // error
	check("monish")  // check for the string
}


```


here the interface{} can be used take any value act as type.