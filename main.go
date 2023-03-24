package main

import (
	"fmt"
	"modules/internal/db"
	"database/sql"
	"log"
)

func DatabaseConnect() *db.Queries{
	dbconn, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/NEWSLETTER")
	if err != nil {
		log.Fatal(err)
	}
	database := db.New(dbconn);
	return database;
}


func main() {
	database := DatabaseConnect();
	fmt.Println(database.GetEmails)
}