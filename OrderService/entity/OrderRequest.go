package entity

type OrderRequest struct {
	UserID int     `json:"user_id"`
	Amount float64 `json:"amount"`
}
