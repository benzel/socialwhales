package storagepq

// import "fmt"

// import "github.com/benzel/socialwhales/app/services/storage"
// import "github.com/benzel/socialwhales/app/models"

// func init() {
// 	var s storagePq
// 	s.Profiles = make(profilesPq)
// 	storage.Register(ALIAS, s)
// }

// type storageMem struct {
// 	Profiles profilesPq
// }

// func (s storagePq) GetProfiles() storage.ProfilesAPI {
// 	return s.Profiles
// }

// type profilesMem map[int64]models.Profile

// func (profiles profilesMem) Create(profile models.Profile) error {
// 	if _, ok := profiles[profile.AccountId]; ok {
// 		return fmt.Errorf("storagemem: profile %d already exists", profile.AccountId)
// 	}
// 	profiles[profile.AccountId] = profile
// 	return nil
// }

// func (profiles profilesMem) Read(id int64) (models.Profile, error) {
// 	profile, ok := profiles[id]
// 	if !ok {
// 		return profile, fmt.Errorf("storagemem: unknown profile %d", id)
// 	}
// 	return profile, nil
// }

// func (profiles profilesMem) Update(profile models.Profile) error {
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
