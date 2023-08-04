package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	defer db.Close();
	if err!=nil{
		fmt.Println("cannot connect to DB");
	}else{
		fmt.Println("connected to DB");
	}

	var email string;
	var password string;

	fmt.Print("enter your email id");
	fmt.Scan(&email);
	fmt.Println();
	fmt.Print("password: ");
	fmt.Scan(&password);


	query:="INSERT INTO validate (email,password) VALUES(?,?)"
	_,err=db.Exec(query,email,password);
	if err!=nil{
		fmt.Println("error cannot insert the values");
	}else{
		fmt.Println("insertion successful..!");
	}
}