package main
import ("fmt")

type Person struct{
	name string
	gender string
}



type Employee struct{
	Person;
	job string
	
}

func NewEmployee(val string,gen string,j string) *Employee{  //constructor
	return &Employee{
		Person:Person{name:val,gender:gen,},
		job:j,
		};
}

func (obj *Employee) print(){                   //struct methods
	fmt.Printf("name : %s gender: %s and job : %s\n",obj.name,obj.gender,obj.job);
}
func main(){
	

  monish:=Employee{
	Person:Person{
		name:"Monish",
		gender:"male",
	},
	job:"DEVOPS",
  }
   dhakshin:=NewEmployee("dhakshin","male","1212");  
 dhakshin.print();   // struct method
  fmt.Println(monish);

}