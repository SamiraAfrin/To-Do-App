package usecase

import (
	"To_Do_App/Task/mocks"
	"To_Do_App/models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateTask(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	nowTime := time.Now().UTC()

	mockTask := models.Task{
		Name:      "task one",
		Status:    "pending",
		Comment:   "nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		tMockTask := mockTask
		tMockTask.ID = 1

		mockTaskRepo.On("Store", mock.Anything, mock.AnythingOfType("*models.Task")).
			Return(nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)

		err := u.Store(context.TODO(), &tMockTask)

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestGetTaskByID(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	nowTime := time.Now().UTC()

	mockTask := models.Task{
		ID:        2,
		Name:      "task one",
		Status:    "pending",
		Comment:   "nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}

	t.Run("success", func(t *testing.T) {
		mockTaskRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).
			Return(&mockTask, nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		task, err := u.GetByID(context.TODO(), 2)

		assert.NoError(t, err)
		assert.Equal(t, mockTask.Name, task.Name)
		assert.Equal(t, mockTask.UserID, task.UserID)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockTaskRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).
			Return(&models.Task{}, errors.New("this task doesn't exist")).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		_, err := u.GetByID(context.TODO(), 1)

		assert.Error(t, err)
		mockTaskRepo.AssertExpectations(t)
	})
}
