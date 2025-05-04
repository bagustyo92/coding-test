package repository

import (
	"errors"
	"sync"
	"wallet/internal/entity"
)

type UserRepository interface {
	GetByID(id string) (*entity.User, error)
	Update(user *entity.User) error
}

type userRepo struct {
	users map[string]*entity.User
	mu    sync.Mutex
}

func NewUserRepository() UserRepository {
	return &userRepo{
		users: map[string]*entity.User{
			"user123": {ID: "user123", Name: "Alice", Balance: 500000},
		},
	}
}

func (r *userRepo) GetByID(id string) (*entity.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	user, exists := r.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (r *userRepo) Update(user *entity.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.ID] = user
	return nil
}
