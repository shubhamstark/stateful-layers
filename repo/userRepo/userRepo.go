package userrepo

import (
	"fmt"
	"sync"

	"github.com/shubhamstark/stateful-layers.git/models/user"
)

type userRepo struct {
	ID   string
	user *user.User
	mu   sync.Mutex
}

// GetUser caches the user on application layer and returns it
func (ur *userRepo) GetUser(returnMemoized bool) user.User {

	if returnMemoized && ur.userCached() {
		return *ur.user
	}

	fmt.Printf("GetUser: fetching user with ID: %v \n", ur.ID)
	result := user.User{Name: "shubham", Age: 5, ID: ur.ID}
	ur.user = &result

	return result

}

func (ur *userRepo) userCached() bool {
	ur.mu.Lock()
	userCached := ur.user != nil
	ur.mu.Unlock()

	return userCached
}

func NewUserRepo(id string) userRepo {
	return userRepo{
		ID: id,
		mu: sync.Mutex{},
	}
}
