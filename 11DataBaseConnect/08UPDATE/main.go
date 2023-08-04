package main


import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main(){

	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	if err!=nil{
		fmt.Println("cannot connect to DB");
	}

	query:="UPDATE table1 SET password=? WHERE email=?"
	
	_,err=db.Exec(query,"8888","monish@gmail.com");
	if err!=nil{
		fmt.Println("cannot update the database");
		return 
	}
	fmt.Println("Successfully updated the database..!");
}