/*
Student Database: 
Build a program that manages student records using structs. Each student should have fields for name, age, and grade.
 Implement features to add new students, update information, and display a student's details.
*/

package main

import "fmt"

type StudentDB struct{
	name string;
	age int;
	grade string;
}

func showDetails(obj StudentDB){
	fmt.Printf("Name: %v\n",obj.name);
	fmt.Printf("age: %v\n",obj.age);
	fmt.Printf("Grade: %v\n",obj.grade);
}
func insert(name string,no int,g string)*StudentDB{
	return &StudentDB{name:name,age:no,grade:g};
}
func update(obj *StudentDB){
	var name string;
	var age int;
	var grade string;
 
	fmt.Println("The Old value is ");

	showDetails(*obj); // why *obj instead of obj ,since showDetails accepts only value not address so i dereference the obj by *obj
	fmt.Print("Enter name: ")
    fmt.Scanf("%v\n", &name) // Notice the "\n" to consume the newline character
    

    fmt.Print("Enter AGE: ")
    fmt.Scanf("%d\n", &age) // Notice the "\n" to consume the newline character

    fmt.Print("Enter grade: ")
    fmt.Scanf("%v\n", &grade) // 
	obj.name=name;
	obj.age=age;
	obj.grade=grade

	fmt.Println("Successfully updated !");
	fmt.Println("New value is :");
	showDetails(*obj)
	fmt.Println()
}

func main(){
	musk:=insert("Elonn musk",40,"O");
	turin:=insert("Alan turin",50,"O+");
	update(musk);
	showDetails(*musk);
	showDetails(*turin);
}