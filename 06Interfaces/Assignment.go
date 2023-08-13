/*

  define a struct student ,get a student details store that using a constructor use interface printer
  interface can contain a pritInfo() function signature,when you want to print details of student you should give 
  print(obj) not obj.print(),the object for each student should named the name of student itself.
*/

package main
import "fmt"


type Student struct{
	stdno string;
	name string;
}
type printer interface{
	printInfo();
}


func(obj Student) printInfo(){
	fmt.Println("Name:",obj.name);
	fmt.Println("student number: ",obj.stdno);
}

func setValue(stdno,name string)Student{
	return Student{stdno: stdno,name: name};
}

func print(p printer){
	p.printInfo()
}
func main(){
	var monish Student;
	monish.name="Monish parameswara"
	monish.stdno="059";
	var name string;
	var no string;
	store:=make(map[string]Student)
	
	fmt.Println("enter name");
	fmt.Scan(&name)
	fmt.Println("enter number");
	fmt.Scan(&no);
	store[name]=setValue(no,name);


	fmt.Println("Enter name to check");
	fmt.Scan(&name);
	if person,ok:=store[name];ok{
        print(person);
	}else{
		fmt.Println("Student not found");
		fmt.Println("Enter number :");
		fmt.Scan(&no);
		store[name]=setValue(no,name)
	}

}