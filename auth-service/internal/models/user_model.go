package models

import "time"

type User struct {
	UserID        string    `json:"user_id"`
	AccountTypeID *string   `json:"account_type_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
