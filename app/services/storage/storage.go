// 2013
// Package storage defines interfaces for persistence service.

package storage

import "github.com/benzel/socialwhales/app/models"

type StorageService interface {
	GetProfiles() ProfilesAPI
	GetAccounts() AccountsAPI
}

type AccountsAPI interface {
	Create(Credentials models.Credentials) (int64, error)
	Check(Credentials models.Credentials) (int64, error)
	UpdateEmail(Credentials models.Credentials, newEmail string) (int64, error)
	UpdatePassword(Credentials models.Credentials, newPassword string) (int64, error)
	Delete(Credentials models.Credentials) (int64, error)
}

type ProfilesAPI interface {
	// TODO(Dyatlov): return more specific errors, e.g. like here
	// http://golang.org/src/pkg/database/sql/driver/driver.go
	Create(profile models.Profile) error
	Read(id int64) (models.Profile, error)
	Update(profile models.Profile) error
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
	if _, ok := storages[alias]; !ok {
		panic("storage: No Storage implementation registered")
	}
	return storages[alias]
}
