package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDBClient() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp(127.0.0.1:3306)/moveinsync")
	if err != nil {
		fmt.Println("There was an error with db")
		panic(err)
	}
	return db
}
