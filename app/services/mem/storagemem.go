// Simple, non-persistent implementation of storage service on top of built-in maps.
// Deprecated
package storagemem

// import "fmt"
// import "code.google.com/p/go.crypto/bcrypt"
// import "github.com/benzel/socialwhales/app/utils"
// import "github.com/benzel/socialwhales/app/models"
// import "github.com/benzel/socialwhales/app/services/storage"

// func init() {
// 	s := &storageMem{
// 		Profiles: make(profilesMem),
// 		Accounts: make(accountsMem),
// 	}
// 	storage.Register("MEM", s)
// }

// type storageMem struct {
// 	Profiles profilesMem
// 	Accounts accountsMem
// }

// func (s storageMem) InitService() error {
// 	return nil
// }

// func (s storageMem) GetProfiles() storage.ProfilesAPI {
// 	return s.Profiles
// }

// func (s storageMem) GetAccounts() storage.AccountsAPI {
// 	return s.Accounts
// }

// type profilesMem map[int64]*models.Profile
// type accountsMem map[string]*models.Account

// func (profiles profilesMem) Create(profile *models.Profile) error {
// 	if _, ok := profiles[profile.AccountId]; ok {
// 		return fmt.Errorf("storagemem: profile %d already exists", profile.AccountId)
// 	}
// 	profiles[profile.AccountId] = profile
// 	return nil
// }

// func (profiles profilesMem) Read(id int64) (*models.Profile, error) {
// 	profile, ok := profiles[id]
// 	if !ok {
// 		return profile, fmt.Errorf("storagemem: unknown profile %d", id)
// 	}
// 	return profile, nil
// }

// func (profiles profilesMem) Update(profile *models.Profile) error {
// 	if _, ok := profiles[profile.AccountId]; !ok {
// 		return fmt.Errorf("storagemem: unknown profile %d", profile.AccountId)
// 	}
// 	profiles[profile.AccountId] = profile
// 	return nil
// }

// func (profiles profilesMem) Delete(id int64) error {
// 	if _, ok := profiles[id]; !ok {
// 		return fmt.Errorf("storagemem: unknown profile %d", id)
// 	}
// 	delete(profiles, id)
// 	return nil
// }

// func (accounts accountsMem) Create(credentials *models.Credentials) (int64, error) {
// 	if _, ok := accounts[credentials.Email]; ok {
// 		return -1, fmt.Errorf("storagemem: account %q already exists", credentials.Email)
// 	}
// 	hPassword, err := utils.HashPassword(credentials.Password)
// 	accountId := int64(len(accounts))
// 	if err != nil {
// 		return -1, fmt.Errorf("storagemem: password generation error")
// 	}
// 	account := &models.Account{
// 		Id:        accountId,
// 		Email:     credentials.Email,
// 		HPassword: hPassword,
// 	}
// 	accounts[account.Email] = account
// 	return accountId, nil
// }

// func (accounts accountsMem) Check(credentials *models.Credentials) (int64, error) {
// 	if account, ok := accounts[credentials.Email]; ok {
// 		err := bcrypt.CompareHashAndPassword([]byte(account.HPassword), []byte(credentials.Password))
// 		if err != nil {
// 			return -1, fmt.Errorf("storagemem: wrong password")
// 		}
// 		return account.Id, nil
// 	}
// 	return -1, fmt.Errorf("storagemem: wrong email")
// }

// func (accounts accountsMem) UpdateEmail(credentials *models.Credentials, newEmail string) (int64, error) {
// 	accountId, err := accounts.Check(credentials)
// 	if err == nil {
// 		account, _ := accounts[credentials.Email]
// 		account.Email = newEmail
// 		return accountId, nil
// 	}
// 	return -1, err
// }

// func (accounts accountsMem) UpdatePassword(credentials *models.Credentials, newPassword string) (int64, error) {
// 	accountId, err := accounts.Check(credentials)
// 	if err == nil {
// 		account, _ := accounts[credentials.Email]
// 		hPassword, err := utils.HashPassword(newPassword)
// 		if err == nil {
// 			account.HPassword = hPassword
// 			return accountId, nil
// 		}
// 		return -1, fmt.Errorf("storagemem: password generation error")
// 	}
// 	return -1, err
// }

// func (accounts accountsMem) Delete(credentials *models.Credentials) (int64, error) {
// 	accountId, err := accounts.Check(credentials)
// 	if err == nil {
// 		delete(accounts, credentials.Email)
// 		return accountId, nil
// 	}
// 	return -1, err
// }
