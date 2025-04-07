package entities

import "database/sql"

type User struct {
	ID           string         `db:"id" json:"id"`
	Email        string         `db:"email" json:"email"`
	PasswordHash string         `db:"password_hash" json:"password_hash"`
	Phone        sql.NullString `db:"phone" json:"phone,omitempty"`
	CreatedAt    float64        `db:"created_at" json:"created_at"`
	UpdatedAt    float64        `db:"updated_at" json:"updated_at"`
}
