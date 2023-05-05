package controller

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
)

func GetBalance(userID int) (float64, error) {
	db, err := ConnectToDB()

	fmt.Println("Connected to Oracle")
	tx, err := db.Begin()

	// logical solve
	var balance float64
	err = db.QueryRow("select BALANCE from MINHTD5.PAYMENTDB where USER_ID = ?", userID).Scan(&balance)
	if err != nil {
		tx.Rollback()
		fmt.Println(err)
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0.0, fmt.Errorf("Failed to commit transaction: %v", err)
	}

	return balance, err
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("godror", "system/oracle@(DESCRIPTION=(ADDRESS_LIST=(ADDRESS=(PROTOCOL=TCP)(HOST=localhost)(PORT=1521)))(CONNECT_DATA=(SERVICE_NAME=orclpdb1)))")
	if err != nil {
		fmt.Println(err)
		return db, nil
	}
	defer db.Close()

	//Thực hiện ping database để đảm bảo kết nối thành công
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
