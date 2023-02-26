package usecase

import (
	"context"
	"time"

	"To_Do_App/User"
	"To_Do_App/models"
)

type userUsecase struct {
	userRepo       User.Repository
	contextTimeout time.Duration
}

func NewUserUsecase(u User.Repository, timeout time.Duration) User.Usecase {

	return &userUsecase{
		userRepo:       u,
		contextTimeout: timeout,
	}
}

// Store new data in the database
func (u *userUsecase) Store(c context.Context, user *models.User) error {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err := u.userRepo.Store(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// Update user info
func (u *userUsecase) Update(c context.Context, user *models.User) error {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	return u.userRepo.Update(ctx, user)

}

// fetch all the task in the task database
func (u *userUsecase) GetAllUser(c context.Context) ([]*models.User, error) {

	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.userRepo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}
