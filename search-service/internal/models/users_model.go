package models

import "time"

type User struct {
	UserID    int       `json:"user_id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
