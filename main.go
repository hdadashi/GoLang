package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	ID    int
	Name  string
	pcode uint
}

// stablishing connection to mysql db
var db, err = sql.Open("mysql", "root:Hossein@tcp(127.0.0.1:3306)/Hossein")

func selectQuery(id int) error {
	results, err := db.Query("SELECT * FROM customers WHERE id = ? and pcode > ?", id, 4)
	if err != nil {
		return err
	}

	defer results.Close()

	fmt.Println("Query Deployed Successfully!")

	// An slice to hold data from returned rows
	var customers []Customer

	// Loop through rows, using Scan to copy data to struct fields
	for results.Next() {
		var cstmr Customer
		results.Scan(&cstmr.ID, &cstmr.Name, &cstmr.pcode)
		customers = append(customers, cstmr)
	}
	if err = results.Err(); err != nil {
		return err
	}
	fmt.Println(customers) //just for testing
	return errors.New("no error found")
}

func addNewTable(Tname string, Tcolumns []string) error {

	// Loop through adding columns
	var Tcolumn string = ""
	for i := 0; i <= len(Tcolumns)-1; i++ {
		Tcolumn = Tcolumn + Tcolumns[i] + ", "
	}
	Tcolumn = string(Tcolumn[0 : len(Tcolumn)-2])
	query := "CREATE TABLE " + Tname + " (" + Tcolumn + ")"
	results, err := db.Query(query)
	if err != nil {
		return err
	}

	defer results.Close()

	fmt.Println("Table Created Successfully!")

	return errors.New("no error found")
}
func main() {

	if err != nil {
		panic(err.Error())
	}
	selectQuery(1)
	var columns = []string{"id int", "name varchar(255)"}
	addNewTable("Users", columns)
	defer db.Close()
	fmt.Println("Success!")
}
