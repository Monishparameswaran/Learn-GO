package main

import (
	"fmt"
	"net/url"
)


func main(){
	partsURL:=&url.URL{
		Scheme:"ht3tps",
		Host:"aws.amazon.com",
		Path:"/free",
	}

	URL:=partsURL.String() // String method is given by url.URL type
	fmt.Println(URL);
}