package models

import "time"

type Account struct {
	Id        int64
	Role      int16
	Email     string
	SignedUp  time.Time
	HPassword []byte
}
