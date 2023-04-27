package controller

func UpdateBalance(balance float64, userID int) (float64, error) {
	db, _ := ConnectToDB()
	_, err := db.Exec("UPDATE service.PAYMENTDB SET BALANCE = ? WHERE USER_ID = ?", balance, userID)
	if err != nil {
		return 0.0, err
	}
	return balance, nil
}
