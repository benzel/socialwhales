package models

import (
    "fmt"
    "time"
    "regexp"
    "github.com/robfig/revel"
)

type User struct {
    UserId         int
    UserEmail      string
    Role           int
    FirstName      string
    Surname        string
    Username       string
    Password       string
    SignedUp       Time
    HashedPassword []byte
}

// func (u *User) String() string {
//     return fmt.Sprintf("User(%s)", u.Username)
// }
// 
// var userRegex = regexp.MustCompile("^\\w*$")
// 
// func (user *User) Validate(v *revel.Validation) {
//     v.Check(user.Username,
//         revel.Required{},
//         revel.MaxSize{15},
//         revel.MinSize{4},
//         revel.Match{userRegex},
//     )
// 
//     ValidatePassword(v, user.Password).
//         Key("user.Password")
// 
//     v.Check(user.Name,
//         revel.Required{},
//         revel.MaxSize{100},
//     )
// }
// 
// func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
//     return v.Check(password,
//         revel.Required{},
//         revel.MaxSize{15},
//         revel.MinSize{5},
//     )
// }