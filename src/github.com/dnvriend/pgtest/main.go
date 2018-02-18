package main

// Package pq is a pure Go Postgres driver for the database/sql package.

import (
	"database/sql"
	"github.com/nu7hatch/gouuid"
	_ "github.com/lib/pq"
	"fmt"
)

func id() string {
	u4, err := uuid.NewV4()
	checkErr(err)
	return u4.String()
}

func main() {
	connStr := "postgres://dnvriend:admin12345@dnvriend-rds.c5mvtqg6mxyp.eu-west-1.rds.amazonaws.com:5432/dnvriend?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	checkErr(err)
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(1)
	fmt.Println("Connected to database")

	stmt, err := db.Prepare("INSERT INTO Transactions (date,year,month,week,day) VALUES($1,$2,$3,$4,$5)")
	checkErr(err)
	for i := 1; i <= 100; i++ {
		insertRecord(stmt)
	}

	rows, err := db.Query("SELECT date, year, month, week, day FROM Transactions")
	checkErr(err)
	for rows.Next() {
		var date string
		var year string
		var month string
		var week string
		var day string
		err = rows.Scan(&date, &year, &month, &week, &day)
		checkErr(err)
		fmt.Printf("%v | %v | %v | %v | %v\n", date, year, month, week, day)
	}
	//fmt.Println("Rows", rows.)
	fmt.Println("Closing connections")
	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertRecord(stmt *sql.Stmt) {
	_, err := stmt.Exec(id(), id(), id(), id(), id())
	checkErr(err)
}
