package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string="https://go.dev/doc/"
func main(){

	data,err:=http.Get(url);
	
	if err!=nil{
		fmt.Println("cannot get the URL");
	}
	defer data.Body.Close();   // its our responsibility to close the request once its done
	content,err:=ioutil.ReadAll(data.Body)
	if err!=nil{
		fmt.Println("cannot read the data");
	}
	//fmt.Println(content);  // its a data that cant be read
	fmt.Println(string(content));
}
