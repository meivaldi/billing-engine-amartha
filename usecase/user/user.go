package usecase

import (
	"github.com/meivaldi/billing-engine/model"
	"github.com/meivaldi/billing-engine/repository"
	uc "github.com/meivaldi/billing-engine/usecase"
)

type UserUsecase struct {
	repositoryDB repository.Repository
}

func New(userRepo repository.Repository) uc.IUserUsecase {
	return &UserUsecase{
		repositoryDB: userRepo,
	}
}

func (uc *UserUsecase) CreateUser(user model.User) (int64, error) {
	id, err := uc.repositoryDB.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (uc *UserUsecase) GetDeliquentUsers() ([]model.User, error) {
	users, err := uc.repositoryDB.GetDeliquentUsers()
	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}
