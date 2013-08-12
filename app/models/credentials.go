package models

import "github.com/robfig/revel"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *Credentials) Validate(v *revel.Validation) {

	v.Required(c.Email)
	v.Email(c.Email)
	v.MaxSize(c.Email, 32)
	v.MinSize(c.Email, 5)

	v.Required(c.Password)
	v.MaxSize(c.Password, 128)
	v.MinSize(c.Password, 6)

}
