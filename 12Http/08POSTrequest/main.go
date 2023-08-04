package main

import (
	"fmt"
	"net/http"
	"strings"
)


func main(){
	// make sure you are running the server , (code in 07POSThandler) code , it is the server which handles this request
	// if not running eexecute this in a new terminal "go run ../07PostHandler/main.go"
	const url string="http://localhost:8080/";

	requestBody:=strings.NewReader(`{
		"name":"monish",
		"country":"India"
	}`)
	
	response,err:=http.Post(url,"application/json",requestBody);
	textRequest,_:=http.Post(url,"text/plain",strings.NewReader("hello monish"));

	if err!=nil{
		fmt.Println("an error has occured");
	}
	fmt.Println("status text :",textRequest.Status);
	fmt.Println("Status code:",response.StatusCode);
}