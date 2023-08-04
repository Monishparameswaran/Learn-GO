package main
import ("fmt")


func main(){
	myslice:=make([]string,10,10);
	fmt.Println(myslice);
	fmt.Println(len(myslice));
	fmt.Println(cap(myslice));
}