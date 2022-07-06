package models

import "time"

type UserModel struct {
	// User unique id
	ID uint
	// Username
	Username string
	// Hashed user password
	Password string

	CreatedAt time.Time
	UpdatedAt time.Time
}
