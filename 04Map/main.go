package main
import ("fmt")

func main(){
	a:=map[string]int{"monish":12,"leo":15,"kio":64};
	
	if _,ok:=a["monish"];ok{
		fmt.Println("it exist");
	}else{
		fmt.Println("it doesnt exist");
	}
}