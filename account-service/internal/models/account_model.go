package models

import "time"

type Account struct {
	AccountID string    `json:"account_id"`
	UserID    string    `json:"user_id" validate:"required"`
	Username  string    `json:"username" validate:"required"`
	FullName  string    `json:"full_name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=6,max=20"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
