package main
import ("fmt")

type printer interface{
	printinfo();
} 
type student struct{
	name string
	sno int
}
type employeer struct{
	id string
	salary int
}
type hr struct{
	age int
	group string
}

func (s student) printinfo(){
	fmt.Println("the name of the student",s.name," ","his sno is:",s.sno)
}

func (e employeer) printinfo(){
	fmt.Println("the name of the employer",e.id," ","salary:",e.salary)
}

func (h hr)printinfo(){
	fmt.Println("Age:",h.age,"group:", h.group);
}

func print(p printer){    // function that print something which takes argument of type interface 
	p.printinfo();
}
func main(){
	var monish student
	monish.name="Monish"
	monish.sno=4572

	var elon employeer
	elon.id="Owner"
	elon.salary=10000
  
	// p:=[]printer{monish,elon};
	// for _,obj:=range p{
	// 	obj.printinfo()
	// }
	print(monish);

}