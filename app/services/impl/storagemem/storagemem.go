package storagemem

import "fmt"

import "github.com/benzel/socialwhales/app/services/storage"
import "github.com/benzel/socialwhales/app/models"

func init() {
    var s storageMem
    s.Users = make(usersMem)
    storage.Register("memory", s)
}

type storageMem struct {
    Users usersMem
}
func (s storageMem) GetUsers() storage.Users {
    return s.Users
}

type usersMem map[string] models.User
func (users usersMem) Create(user models.User) error {
    if _, ok := users[user.Login]; ok {
        return fmt.Errorf("storagemem: user %q already exists", user.Login)
    }
    users[user.Login] = user
    return nil
}
func (users usersMem) Read(login string) (models.User, error) {
    user, ok := users[login]
    if !ok {
        return user, fmt.Errorf("storagemem: unknown user %q", login)
    }
    return user, nil
}
func (users usersMem) Update(user models.User) error {
    if _, ok := users[user.Login]; !ok {
        return fmt.Errorf("storagemem: unknown user %q", user.Login)
    }
    users[user.Login] = user
    return nil
}
func (users usersMem) Delete(login string) error {
    if _, ok := users[login]; !ok {
        return fmt.Errorf("storagemem: unknown user %q", login)
    }
    delete(users, login)
    return nil
}