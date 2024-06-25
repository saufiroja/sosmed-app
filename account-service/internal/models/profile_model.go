package models

import "time"

type Profile struct {
	ProfileID   string    `json:"profile_id"`
	UserID      string    `json:"user_id"`
	Avatar      string    `json:"avatar"`
	Bio         string    `json:"bio"`
	Location    string    `json:"location"`
	Website     string    `json:"website"`
	BirthDate   time.Time `json:"birth_date"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
