// 2013
// Package storage defines interfaces for persistence service.

package storage

import "github.com/benzel/socialwhales/app/models"

type Storage interface {
	GetUsers() Users
}

type Users interface {
	// TODO(Dyatlov): return more specific errors, e.g. like here
	// http://golang.org/src/pkg/database/sql/driver/driver.go
	Create(user models.User) error
	Read(login string) (models.User, error)
	Update(user models.User) error
	Delete(login string) error
}

var storages = make(map[string]Storage)

func Register(name string, storage Storage) {
	if storage == nil {
		panic("storage: Register storage is nil")
	}
	if _, ok := storages[name]; ok {
		panic("storage: Register called twice for storage implementation " + name)
	}
	storages[name] = storage
}

func GetStorage(name string) Storage {
	if _, ok := storages[name]; !ok {
		panic("storage: No Storage implementation registered")
	}
	return storages[name]
}
