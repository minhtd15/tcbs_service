package entity

type Payment struct {
	UserID  int     `json:"user_id"`
	Balance float64 `json:"balance"`
}
