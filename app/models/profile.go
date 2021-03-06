package models

type Profile struct {
	AccountId  int64  `json:"account_id,db:"account_id"`
	ScreenName string `json:"screen_name,omitempty",db:"screen_name"`
	FirstName  string `json:"first_name,omitempty",db:"first_name"`
	Surname    string `json:"surname,omitempty",db:"surname"`
	Email      string `json:"email,omitempty",db:"email"`
}

// Constructs profile for provided account data, can be used when registering a new user.
func InitialProfile(account *Account) *Profile {
	return &Profile{
		AccountId:  account.Id,
		ScreenName: "",
		FirstName:  "",
		Surname:    "",
		Email:      "",
	}
}
