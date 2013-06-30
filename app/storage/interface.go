/**
 * This file declares the interfaces used to access persistent storage.
 */

package storage

import (
	"socialwhales/app/models"
)

type IAccountAPI interface {
	Create(Credentials* models.AccountCredentials) (error, int64)
	Check(Credentials* models.AccountCredentials) (error, int64)
	EmailExists(Email string) (error, int64)
	UpdateEmail(Credentials* models.AccountCredentials, NewEmail string) (error, int64)
	UpdatePassword(Credentials* models.AccountCredentials, NewPassword string) (error, int64)
	Delete(Credentials* models.AccountCredentials) (error, int64)
}

type IStorageService interface {
	GetAccountAPI() *IAccountAPI
}


// var STORAGE *IStorageService