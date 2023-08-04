package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // while encoding the json will take coursename instead of Name
	email    string
	Password string `json:"-"` // giving this will omit the password field when json encoding and decoding it will not be transferred
	Price    int
	Type     string
}

func EncodeJson() {
	// create a object
	web := &course{"web", "monish@gmail.com", "monish122", 4355, "online"}
	dev := &course{"dev", "opensource@gmail.com", "7854", 0, "offline"}

	// it is simple converting a struct data into json 

	//Marshal is  a encoding into json 
	DataJson, err := json.Marshal(web)
	DevJson, err := json.MarshalIndent(dev, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", DataJson)
	fmt.Println();

	fmt.Println("the following is the json data that is indented using .MarshalIndent");

	fmt.Printf("\n%s\n", DevJson)
}
func main() {
	EncodeJson()
}
