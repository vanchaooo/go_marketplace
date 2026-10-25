package repository

import (
	"math"
	"strings"
	"slices"
	"github.com/vanchaooo/go-marketplace/internal/model"
)

type UserRepository struct {
	users map[int64]*model.User
	nextID int64
}

func NewUserRepository() *UserRepository {
	return &UserRepository {
		users: make(map[int64]*model.User),
		nextID: 1,
	}
}

func (u *UserRepository) Add(user *model.User) (int64, bool) {
	if u.nextID <= 0 || u.nextID == math.MaxInt64 {
		return 0, false
	}
 
	user.ID = u.nextID
	u.users[user.ID] = user
	u.nextID++
	return user.ID, true
}

func (u *UserRepository) Get(userID int64) (*model.User, bool) {
	user, ok := u.users[userID]
	return user, ok
}

func (u *UserRepository) GetEveryoneByName(name string) []int64 {
	name = strings.ToLower(name)
	result := []int64{}

	for id, user := range u.users {
		if strings.Contains(strings.ToLower(user.Name), name) {
			result = append(result, id)
		}
	}
	slices.Sort(result)
	return result
}

func (u *UserRepository) Delete(userID int64) {
	delete(u.users, userID)
}

func (u *UserRepository) Exists(userID int64) bool {
	_, ok := u.users[userID]
	return ok
}