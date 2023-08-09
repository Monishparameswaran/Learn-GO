package main

import "fmt"

/* What is anonymous struct
     These are used if you are not going to use the struct not more than one time , in case of developing a software application
	 we may have n number struct,so we have to give appropriate name for each ,to avoid naming a on-off ,one used or short lived 
	 struct ,we go for anonymous struct these are just like other struct but can be used only onetime,since we are declaring at that time we
	 needed and not giving any names to it. 
*/
func main(){

	company:=struct{
		name string
		location string
	}{
		name:"Tesla",
		location:"india",
	}

	fmt.Println(company); 

}