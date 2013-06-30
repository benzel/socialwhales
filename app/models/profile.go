package models

import (
	// "database/sql"
	// _ "github.com/bmizerany/pq"
	"github.com/robfig/revel"
	"regexp"
	// "time"
)

type Profile struct {
	AccountId  int64  `json:"account_id"`
	ScreenName string `json:"screen_name,omitempty"`
	FirstName  string `json:"first_name,omitempty"`
	Surname    string `json:"surname,omitempty"`
}

// func (p *Profik) Get(Email string, Password string) {
// 	db, err := sql.Open("postgres", "user=zvm password=39572685 dbname=whale sslmode=disable")
// 	if err != nil {
// 		// TODO
// 		revel.INFO.Println(err)
// 	}
// 	revel.INFO.Println("Email: " + Email)
// 	revel.INFO.Println("Password: " + Password)
// 	results, err := db.Query(`SELECT id,role,email,screen_name,first_name,surname
//                               FROM t_users
//                               WHERE email=$1 AND hashed_password=$2;`, &Email, &Password)
// 	if err != nil {
// 		// TODO
// 		revel.INFO.Println(err)
// 	}
// 	for results.Next() {
// 		var id int
// 		var name string
// 		revel.INFO.Println("Name: " + name)
// 		err = results.Scan(&id, &name)
// 	}
// }
//
// func (u *User) Create(c *UserCredentials) {
//
// }

func (p *Profile) Validate(v *revel.Validation) {
	var wordRegex = regexp.MustCompile("^\\w*$")

	v.Check(p.ScreenName,
		revel.MaxSize{16},
		revel.MinSize{4},
		revel.Match{wordRegex},
	)

	v.Check(p.FirstName,
		revel.MaxSize{16},
		revel.MinSize{1},
		revel.Match{wordRegex},
	)

	v.Check(p.Surname,
		revel.MaxSize{16},
		revel.MinSize{1},
		revel.Match{wordRegex},
	)
}

// type UserToken struct {
// 	AccessKey      string `json:"key"`
// 	AccessPassword string `json:"secret"`
// }
