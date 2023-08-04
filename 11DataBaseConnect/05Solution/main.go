package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


func signup(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");

	if err!=nil{
		fmt.Println("cannot connect to the DB");
	}

	query:="INSERT INTO validate (email,password) VALUES(?,?)"

	fmt.Println("--------------SignUp--------------------------")
	var name string;
	var fpass string;
	var lpass string;
	fmt.Print("enter name: ");
	fmt.Scan(&name);
	fmt.Print("enter password: ");
	fmt.Scan(&fpass);
	fmt.Println("renter your password");
	fmt.Scan(&lpass);
	if(fpass!=lpass){
		fmt.Println("PassWord-Mismatch, cannot sign up try again");
	}else{
		_,err=db.Exec(query,name,lpass);
	if err!=nil{
		fmt.Println("cannot insert the value");
	}else{
		fmt.Println("SignUP successful....!");
		fmt.Println("--- Try Login---------");
		time.Sleep(time.Second * 2);
		fmt.Println();
		route(2);
	}
	}
	
}

func userExist(name string)bool{
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	if err!=nil{
		fmt.Println("cannot connect to DB in userExist");
	}

	query:="SELECT * FROM validate"

	row,err:=db.Query(query);
	if err!=nil{
		fmt.Println("Invalid query");
	}

	for row.Next(){
		var user string;
		var pass string;
		row.Scan(&user,&pass);
		if(user==name){
			
			return true
		}
	}
	return false
}
func login(){

	var email string;
	var password string;
	fmt.Println("-----------------Login----------------------");
	fmt.Print("enter email: ");
	fmt.Scan(&email)
	fmt.Print("enter password: ");
	fmt.Scan(&password);
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	if err!=nil{
		fmt.Println("cannot able to connect to DB");
	}
	query:="SELECT * FROM validate";
	row,err:=db.Query(query);

	if(!userExist(email)){
		fmt.Println("UserNotExist...");
		fmt.Println("You will be redirected to signup page...");
		time.Sleep(time.Second * 3);
		fmt.Println()
		route(1);
	}else{
		for row.Next(){
		var name string;
		var pass string;
		row.Scan(&name,&pass);
		if(name==email){
			if(pass!=password){
				fmt.Println("Password does not match, reenter it..");
			}else{
				fmt.Println("Welcome user, login successful...!");
			}
		}
		
	}
}
}

func route(num int){
	switch(num){
	case 1:
		signup();
	case 2:
		login();
	}
}
func main(){

	

	fmt.Println("Enter 1 Signup");
	fmt.Println("Enter 2 Login in");

	var num int;

	fmt.Scan(&num);

	route(num);
}