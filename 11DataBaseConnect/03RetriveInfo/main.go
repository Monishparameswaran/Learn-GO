package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	db,err:=sql.Open("mysql","root:@tcp(localhost:3306)/storeDB");
	defer db.Close();
	if err!=nil{
		fmt.Println("cannot connect to the DB");
	}else{
		fmt.Println("connected to DB...");
	}

	query:="SELECT * FROM validate"
    row,err:=db.Query(query)
	if err!=nil{
		fmt.Println("invalid queury");
	}

	for row.Next(){	
		var email string;
		var pass string;

		err=row.Scan(&email,&pass);
		if err!=nil{fmt.Println("cannot read DB ",err);}
		fmt.Println("email",email,"password: ",pass);

	}

	if row.Err()!=nil{
		fmt.Println(err);
	}
}