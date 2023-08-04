package main
import (
	"fmt"
	"database/sql"
     _"github.com/go-sql-driver/mysql"
)


func main(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	if err!=nil{
		fmt.Println("cannot connect to DB");
	}
	var email string;
	fmt.Print("enter your mail to update tp password");
	fmt.Scan(&email);
	var pass string;
	fmt.Print("enter pass to update:");
	fmt.Scan(&pass);
	query:="UPDATE table1 SET password=? WHERE email=?"
	_,err=db.Exec(query,pass,email);
	if err!=nil{
		fmt.Println("Cannot update the password");
		return 
	}

	fmt.Println("Updated successfully");
}