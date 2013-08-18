package models

import "time"
import "github.com/robfig/revel"
import uuid "github.com/nu7hatch/gouuid"
import "github.com/benzel/socialwhales/app/utils"

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

func (c *Credentials) Account() (*Account, error) {
	hPassword, err := utils.HashPassword(c.Password)
	if err != nil {
		return nil, err
	}
	tokenObj, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	account := &Account{
		Role:       0,
		Status:     0,
		Email:      c.Email,
		HPassword:  hPassword,
		Token:      tokenObj.String(),
		SignedUpAt: now,
		UpdatedAt:  now,
	}
	return account, nil
}
