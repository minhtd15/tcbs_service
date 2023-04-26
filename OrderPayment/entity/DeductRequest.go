package entity

type DeductRequest struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}
