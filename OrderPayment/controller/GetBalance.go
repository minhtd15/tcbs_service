package controller

import (
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetBalance(userID int) float64 {
	// db connector
	db, err := connectToDB()

	// logical solve
	var rs float64
	row := db.QueryRow("SELECT balance FROM MINHTD5.PAYMENTDB WHERE USER_ID = ?", userID)
	err = row.Scan(&rs)
	if err != nil {
		return 0.0
	}
	return rs
}

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("godror", "system/oracle@localhost/orclpdb1")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	fmt.Println("Connected to Oracle")
	return db, err
}