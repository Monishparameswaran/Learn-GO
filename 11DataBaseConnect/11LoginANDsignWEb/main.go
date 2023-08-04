package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	
)

const url string="https://www.lco.dev"
func main(){

	response,err:=http.Get(url)
	
	if err!=nil{
		fmt.Println("cannot get the url");
	}
	defer response.Body.Close()
	databytes,err:=ioutil.ReadAll(response.Body)
	if err!=nil{
		fmt.Println("cannot read the package");
	}
	content:=string(databytes)
	fmt.Println(content)

}