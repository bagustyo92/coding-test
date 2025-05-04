package service

import (
	"errors"
	"wallet/internal/entity"
	"wallet/internal/repository"
)

type DisbursementService interface {
	Disburse(userID string, amount float64) (*entity.User, error)
	GetUserByID(userID string) (*entity.User, error)
}

type disbursementService struct {
	userRepo repository.UserRepository
}

func NewDisbursementService(userRepo repository.UserRepository) DisbursementService {
	return &disbursementService{userRepo: userRepo}
}

func (s *disbursementService) Disburse(userID string, amount float64) (*entity.User, error) {
	if amount <= 0 {
		return nil, errors.New("invalid disbursement amount")
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user.Balance < amount {
		return nil, errors.New("insufficient balance")
	}

	user.Balance -= amount
	s.userRepo.Update(user)
	return user, nil
}

func (s *disbursementService) GetUserByID(userID string) (*entity.User, error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
