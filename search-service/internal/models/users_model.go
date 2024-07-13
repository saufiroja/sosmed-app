package models

import "time"

type User struct {
	UserID    string    `json:"user_id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
