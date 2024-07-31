package db

import "time"

type Entry struct {
	ID        int64     `json:"id`
	AccountID int64     `json:"accoount_id"`
	Ammount   int64     `json: "ammount"`
	CreatedAt time.Time `json: "created_at"`
}
