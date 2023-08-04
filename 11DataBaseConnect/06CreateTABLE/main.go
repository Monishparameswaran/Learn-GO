package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	if err!=nil{
		fmt.Println("cannot connect to the DB");
	}
	defer db.Close();
	query:=`CREATE TABLE IF NOT EXISTS table1(
		id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(20) NOT NULL,
		password VARCHAR(20) NOT NULL
	);`
	_,err=db.Exec(query);
	if err!=nil{
		fmt.Println("wrong query ",err);
		return;	
	}
	fmt.Println("Created TABLE");
}