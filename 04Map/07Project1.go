package main

/*
Online Shopping Cart: Build a simple online shopping cart system 
where users can add and remove items from their cart. Use a map to store the items and their quantities.
*/

import "fmt"

var cart=make(map[string]int);

func addTocart(item string,q int){
	cart[item]=q;
}
func listCart(){
	fmt.Println("item"," "," quatity");
	for i,v:=range cart{
		
		fmt.Println(i," ",v);
	}
}

func removeFromCart(item string){
	_,ok:=cart[item];
	if ok{
	  fmt.Printf("%v removed from cart !",item);
	  delete(cart,item);	
	}else{
		fmt.Println("the entered item is not in cart !");
	}
	
}
func input(){
	var item string;
	var quantities int;
	fmt.Print("Enter item: ");
	fmt.Scanf("%v",&item);
	fmt.Println();
	fmt.Print("Enter number: ");
	fmt.Scanf("%v",&quantities);
	addTocart(item,quantities)
}
func main(){

	input();
	addTocart("MacbookPro",3)
	addTocart("keyboard",4);
	addTocart("mouse ",5);
	listCart();
	removeFromCart("MacbookPro");
	listCart()
}