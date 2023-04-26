package controller

func UpdateBalance(balance float64, userID int) error {
	db, _ := ConnectToDB()
	_, err := db.Exec("UPDATE MINHTD5.PAYMENTDB SET BALANCE = ? WHERE USER_ID = ?", balance, userID)
	if err != nil {
		return err
	}
	return nil
}
