package storagepq

import "fmt"
import "database/sql"
import "code.google.com/p/go.crypto/bcrypt"
import "github.com/coopernurse/gorp"
import "github.com/robfig/revel"
import "github.com/benzel/socialwhales/app/models"
import "github.com/benzel/socialwhales/app/services/storage"

/////////////////////////////////////////////////////////////////////

func init() {
	storage.Register("PQ", &storagePq{})
}

/////////////////////////////////////////////////////////////////////

type storagePq struct {
	Configured bool
	Profiles   *profilesPq
	Accounts   *accountsPq
	Db         *sql.DB
	DbSpec     string
	DbDriver   string
	DbMap      *gorp.DbMap
}

type profilesPq struct {
	ProfilesTb *gorp.TableMap
	Storage    *storagePq
}

type accountsPq struct {
	AccountsTb *gorp.TableMap
	Storage    *storagePq
}

/////////////////////////////////////////////////////////////////////

// Reads DB configs and initialized backend.
func (s *storagePq) InitService() error {
	var found bool
	var err error

	s.DbDriver, found = revel.Config.String("db.driver")
	if !found {
		revel.ERROR.Fatal("No db.driver found.")
	}

	s.DbSpec, found = revel.Config.String("db.spec")
	if !found {
		revel.ERROR.Fatal("No db.spec found.")
	}

	revel.INFO.Println("Initializing database connection (" + s.DbDriver + ")")
	s.Db, err = sql.Open(s.DbDriver, s.DbSpec)
	if err != nil {
		revel.ERROR.Println(err.Error())
	}

	s.DbMap = &gorp.DbMap{Db: s.Db, Dialect: gorp.PostgresDialect{}}
	if s.DbMap == nil {
		revel.ERROR.Println("Cannot map database " + s.DbDriver)
	}

	s.Accounts = &accountsPq{
		Storage:    s,
		AccountsTb: s.DbMap.AddTableWithName(models.Account{}, "t_accounts").SetKeys(true, "id"),
	}
	s.Profiles = &profilesPq{
		Storage:    s,
		ProfilesTb: s.DbMap.AddTableWithName(models.Profile{}, "t_profiles"),
	}
	s.Configured = true
	revel.INFO.Println(s)
	return nil
}

func (s *storagePq) GetProfiles() storage.ProfilesAPI {
	return s.Profiles
}

/////////////////////////////////////////////////////////////////////

func (s *storagePq) GetAccounts() storage.AccountsAPI {
	return s.Accounts
}

/////////////////////////////////////////////////////////////////////

func (profiles *profilesPq) Create(profile *models.Profile) error {
	return nil
}

/////////////////////////////////////////////////////////////////////

func (profiles *profilesPq) Read(id int64, profile *models.Profile) error {
	return nil
}

/////////////////////////////////////////////////////////////////////

func (profiles *profilesPq) Update(profile *models.Profile) error {
	return nil
}

/////////////////////////////////////////////////////////////////////

func (profiles *profilesPq) Delete(id int64) error {
	return nil
}

/////////////////////////////////////////////////////////////////////

func (accounts *accountsPq) Create(credentials *models.Credentials) (*models.Account, error) {
	account, err := credentials.Account()
	if err != nil {
		return nil, err
	}
	if err = accounts.Storage.DbMap.Insert(account); err != nil {
		return nil, err
	}
	return account, nil
}

/////////////////////////////////////////////////////////////////////

func (accounts *accountsPq) Read(credentials *models.Credentials) (*models.Account, error) {
	var accountList []*models.Account
	query := "SELECT id,email,hpassword,token FROM t_accounts WHERE email=$1"
	_, err := accounts.Storage.DbMap.Select(&accountList, query, credentials.Email)
	if llen := len(accountList); llen != 1 {
		return nil, fmt.Errorf("storagemem: Got %d users instead of 1", llen)
	}
	account := accountList[0]
	err = bcrypt.CompareHashAndPassword([]byte(account.HPassword), []byte(credentials.Password))
	if err != nil {
		return nil, fmt.Errorf("storagemem: wrong password")
	}
	return account, nil
}

/////////////////////////////////////////////////////////////////////

func (accounts *accountsPq) UpdateEmail(credentials *models.Credentials, newEmail string) error {
	return nil
}

/////////////////////////////////////////////////////////////////////

func (accounts *accountsPq) UpdatePassword(credentials *models.Credentials, newPassword string) error {
	return nil
}

/////////////////////////////////////////////////////////////////////

func (accounts *accountsPq) Delete(credentials *models.Credentials) error {
	return nil
}
