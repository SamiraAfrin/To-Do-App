package usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"To_Do_App/User/mocks"
	"To_Do_App/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {

	mockUserRepo := new(mocks.Repository)

	mockUser := models.User{
		Name: "First User",
	}

	t.Run("success", func(t *testing.T) {
		tMockUser := mockUser
		tMockUser.ID = 1

		mockUserRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.User")).
			Return(nil).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)

		err := u.Store(context.TODO(), &tMockUser)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestGetAllUser(t *testing.T) {

	mockUserRepo := new(mocks.Repository)
	mockUser := &models.User{}
	mockListUser := make([]*models.User, 0)
	mockListUser = append(mockListUser, mockUser)

	t.Run("success", func(t *testing.T) {
		mockUserRepo.On("GetAllUser", mock.Anything).
			Return(mockListUser, nil).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		_, err := u.GetAllUser(context.TODO())

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockUserRepo.On("GetAllUser", mock.Anything).
			Return(mockListUser, errors.New("user table doesn't have any data")).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		list, err := u.GetAllUser(context.TODO())

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockUserRepo.AssertExpectations(t)
	})
}

func TestUserUpdate(t *testing.T) {

	mockUserRepo := new(mocks.Repository)
	mockUser := &models.User{
		Name: "Update name",
	}

	t.Run("success", func(t *testing.T) {

		mockUserRepo.On("Update", mock.Anything, mockUser).
			Return(nil).Once()

		u := NewUserUsecase(mockUserRepo, time.Second*2)
		err := u.Update(context.TODO(), mockUser)

		assert.NoError(t, err)
		mockUserRepo.AssertExpectations(t)
	})
}
