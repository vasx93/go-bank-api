package model

import "time"

type Transfer struct {
	ID            int64     `json:"ID"`
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Ammount       int64     `json: "ammount"`
	CreatedAt     time.Time `json: "created_at"`
}
