package main

import (
	"fmt"
	"errors"
)

func Divide(a,b int) (int,error){
	if(b!=0){
		return a/b,nil;
	}else{
	    return 0,errors.New("Cannot divide by zero !");
	}
}
func main(){
	var a,b int;
	fmt.Print("Enter value of a: ");
	fmt.Scanf("%d",&a);
	fmt.Print("ENter value of b: ");
	fmt.Scanf("%d",&b);

	ans,err:=Divide(a,b);
	// if err!=nil{
	// 	fmt.Println(err);
	// }
	fmt.Println(ans," ",err);
}