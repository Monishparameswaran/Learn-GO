package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // while encoding the json will take coursename instead of Name
	email    string
	Password string  `json:"-"`// giving this will omit the password field when json encoding and decoding it will not be transferred
	Price    int
	Type     string
}

var DecodedJsonData course;

func DecodeJson(){

	EncodedJson:=[]byte(
		`{
			"coursename": "dev",
			"email":"monisjh@gmail.com",
			"Price": 0,
			"Type": "offline"
	}`)

	checkvalid:=json.Valid(EncodedJson)
	if checkvalid{

	    json.Unmarshal(EncodedJson,&DecodedJsonData)
		fmt.Println("JSON is valid ");
		fmt.Printf("%#v\n",DecodedJsonData); // json is decoded into structure data
	}else{
		fmt.Println("it is not a valid json data");
	}

}


func main() {
	DecodeJson(); 
}
