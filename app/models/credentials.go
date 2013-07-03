package models

import (
	"github.com/robfig/revel"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) Validate(v *revel.Validation) {
	v.Check(c.Email,
		revel.Required{},
		revel.MaxSize{32},
		revel.MinSize{5},
		revel.Email{},
	)
	v.Check(c.Password,
		revel.Required{},
		revel.MaxSize{128},
		revel.MinSize{6},
		revel.Email{},
	)
}
