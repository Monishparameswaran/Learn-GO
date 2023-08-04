package main

import ("fmt")

func addArr(a,b [3]int) []int{
	c:=[3]int{};
	for i:=0;i<len(a);i++{
		c[i]=a[i]+b[i];
	}
	return c[0:3]
}
func main(){
	var a =[3]int{1,2,3};
	var b =[3]int{7,2,3};

	fmt.Println(addArr(a,b));

}