package models

type Profile struct {
	AccountId  int64  `json:"account_id"`
	ScreenName string `json:"screen_name,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	Surname    string `json:"surname,omitempty"`
	Email      string `json:"email,omitempty"`
}
