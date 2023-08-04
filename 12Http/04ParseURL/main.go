package main

import (
	"fmt"
	"net/url"
)

const URL string="https://example.com:8080/path?key1=value1&key2=value2"
func main(){
	result,err:=url.Parse(URL);

	if err!=nil{
		fmt.Println("cannot process the URL");
	}

	fmt.Println("scheme: ",result.Scheme);
	fmt.Println("Path: ",result.Path);
	fmt.Println("Host: ",result.Hostname());
	fmt.Println("Port: ",result.Port());
	fmt.Println("Raw query: ",result.RawQuery);
}