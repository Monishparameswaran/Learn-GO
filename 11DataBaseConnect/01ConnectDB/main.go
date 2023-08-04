package main
import (
	"fmt"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)


func main(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB") // establish a connection

	/*  format
	  db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/database_name")
	*/
	defer db.Close();
	if err!=nil{
		fmt.Println("cannot able to connect to the server");
	}else{
		fmt.Println("connected to database..!");
	}
	
}