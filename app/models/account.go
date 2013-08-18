package models

import "time"

type Account struct {
	Id         int64     `json:"id" db:"id"`
	Role       int16     `db:"role"`
	Status     int16     `db:"status"`
	Email      string    `json:"email" db:"email"`
	HPassword  string    `json:- db:"hpassword"`
	Token      string    `json:"token" db:"token"`
	SignedUpAt time.Time `db:"signedup_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}
