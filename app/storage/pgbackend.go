package storage

import (
	// "errors"
	"github.com/robfig/revel"
	"socialwhales/app/models"
	_ "database/sql"
	"database/sql/driver"
	"github.com/bmizerany/pq"
)


type AccountAPI struct {
	conn *driver.Conn
}


func NewAccountAPI(db *driver.Conn) *AccountAPI{
	conn, err := pq.Open("user=whale password=password host=localhost port=5432 dbname=whale sslmode=disable")
	if err != nil {
		revel.ERROR.Print("Cannot create connection to PostgreSQL")
	}
	return &AccountAPI{
		conn: &conn,
	}
}

func (a AccountAPI) Create(Credentials* models.AccountCredentials) (error, int64) {
	return nil, -1
}

func (a AccountAPI) Check(Credentials* models.AccountCredentials) (error, int64) {
	return nil, -1
}

func (a AccountAPI) EmailExists(Email string) (error, int64) {
	return nil, -1
}

func (a AccountAPI) UpdateEmail(Credentials* models.AccountCredentials, NewEmail string) (error, int64) {
	return nil, -1
}

func (a AccountAPI) UpdatePassword(Credentials* models.AccountCredentials, NewPassword string) (error, int64) {
	return nil, -1
}

func (a AccountAPI) Delete(Credentials* models.AccountCredentials) (error, int64) {
	return nil, -1
}


type PqStorage struct {
	accountAPI *AccountAPI;
}

func NewPqStorage() *PqStorage {
	return &PqStorage{
		accountAPI: &AccountAPI{},
	}
}

func (s PqStorage) GetAccountAPI() *IAccountAPI {
	return s.accountAPI
}