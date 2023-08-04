package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostHandler(w http.ResponseWriter,r *http.Request){
	data,err:=ioutil.ReadAll(r.Body);
	if err!=nil{
	    http.Error(w,err.Error(),http.StatusInternalServerError);
	}
	fmt.Println(string(data));
	w.WriteHeader(http.StatusOK);
}
func main(){

	http.HandleFunc("/",PostHandler)
	http.ListenAndServe(":8080",nil)
}