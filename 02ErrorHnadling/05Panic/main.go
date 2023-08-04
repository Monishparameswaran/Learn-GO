package main

import ("fmt")

func checkName(name string){
	if (name!="Monish"){
		// I m panic , when i executed i terminate the program !
		panic("enter a valid name !");
	}
}
func main(){
	var name string;
	fmt.Print("please enter name:");
	fmt.Scanf("%s",&name);

	checkName(name);
}