package models

import "time"

type Account struct {
	Id        int64     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	HPassword string    `json:- db:"hpassword"`
	Token     string    `json:"token" db:"token"`
	Role      int16     `db:-`
	Status    int16     `db:-`
	SignedUp  time.Time `db:-`
	UpdatedAt time.Time `db:-`
}
