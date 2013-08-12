// 2013
// Package storage defines interfaces for persistence service.

package storage

import "github.com/robfig/revel"
import "github.com/benzel/socialwhales/app/models"

var Storage StorageService

type StorageService interface {
	InitService() error
	GetProfiles() ProfilesAPI
	GetAccounts() AccountsAPI
}

type AccountsAPI interface {
	Create(credentials *models.Credentials) (*models.Account, error)
	Read(Credentials *models.Credentials) (*models.Account, error)
	UpdateEmail(Credentials *models.Credentials, newEmail string) error
	UpdatePassword(Credentials *models.Credentials, newPassword string) error
	Delete(Credentials *models.Credentials) error
}

type ProfilesAPI interface {
	// TODO(Dyatlov): return more specific errors, e.g. like here
	// http://golang.org/src/pkg/database/sql/driver/driver.go
	Create(profile *models.Profile) error
	Read(id int64, profile *models.Profile) error
	Update(profile *models.Profile) error
	Delete(id int64) error
}

var storages = make(map[string]StorageService)

func Register(alias string, storage StorageService) {
	if storage == nil {
		panic("storage: Register storage is nil")
	}
	if _, ok := storages[alias]; ok {
		panic("storage: Register called twice for storage implementation " + alias)
	}
	storages[alias] = storage
}

func GetStorage(alias string) StorageService {
	revel.INFO.Println(storages)
	if _, ok := storages[alias]; !ok {
		revel.ERROR.Fatal("storage: No Storage implementation registered")
	}
	return storages[alias]
}

func InitServices() {
	for serviceAlias, service := range storages {
		revel.INFO.Println("Init service: " + serviceAlias)
		service.InitService()
		revel.INFO.Println("Should be inited")
		revel.INFO.Println(service)
	}
}
