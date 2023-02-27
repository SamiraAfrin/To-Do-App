package usecase

import (
	"To_Do_App/Task"
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
	//nowTime := time.Now().UTC()

	mockTask := models.Task{
		ID: 2,
		//Name:      "task one",
		//Status:    "pending",
		//Comment:   "nai",
		//UpdatedAt: &nowTime,
		//CreatedAt: &nowTime,
		UserID: 1,
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

func TestGetTaskByUserID(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	//nowTime := time.Now().UTC()

	mockTask := &models.Task{
		ID: 2,
		//Name:      "task one",
		//Status:    "pending",
		//Comment:   "nai",
		//UpdatedAt: &nowTime,
		//CreatedAt: &nowTime,
		UserID: 1,
	}
	mockListTask := make([]*models.Task, 0)
	mockListTask = append(mockListTask, mockTask)
	t.Run("success", func(t *testing.T) {
		mockTaskRepo.On("GetByUserID", mock.Anything, mock.AnythingOfType("int64")).
			Return(mockListTask, nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		tasks, err := u.GetByUserID(context.TODO(), 2)

		assert.NoError(t, err)
		for _, task := range tasks {
			assert.Equal(t, mockTask.UserID, task.UserID)
		}

		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockTaskRepo.On("GetByUserID", mock.Anything, mock.AnythingOfType("int64")).
			Return(mockListTask, errors.New("this user doesn't have any task")).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		list, err := u.GetByUserID(context.TODO(), 2)

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockTaskRepo.AssertExpectations(t)
	})
}
func TestGetAllTask(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	mockTask := &models.Task{}
	mockListTask := make([]*models.Task, 0)
	mockListTask = append(mockListTask, mockTask)

	t.Run("success", func(t *testing.T) {
		mockTaskRepo.On("GetAllTask", mock.Anything).
			Return(mockListTask, nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		_, err := u.GetAllTask(context.TODO())

		assert.NoError(t, err)
		mockTaskRepo.AssertExpectations(t)
	})

	t.Run("not-found", func(t *testing.T) {
		mockTaskRepo.On("GetAllTask", mock.Anything).
			Return(mockListTask, errors.New("task table doesn't have any data")).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		list, err := u.GetAllTask(context.TODO())

		assert.Error(t, err)
		assert.Len(t, list, 0)
		mockTaskRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	nowTime := time.Now().UTC()

	mockTask := &models.Task{
		ID:        2,
		Name:      "task one",
		Status:    "pending",
		Comment:   "nai",
		UpdatedAt: &nowTime,
		//CreatedAt: &nowTime,
		UserID: 1,
	}

	t.Run("success", func(t *testing.T) {

		mockTaskRepo.On("Update", mock.Anything, mockTask).
			Return(nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		err := u.Update(context.TODO(), mockTask)

		assert.NoError(t, err)
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

func TestTaskDone(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	nowTime := time.Now().UTC()

	modelTask := models.Task{}
	mockTask := &Task.TaskPatchReq{
		Status:    "pending",
		UpdatedAt: &nowTime,
	}

	t.Run("success", func(t *testing.T) {

		mockTaskRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).
			Return(&modelTask, nil).Once()
		mockTaskRepo.On("UpdateDone", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("*models.Task")).
			Return(nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		err := u.UpdateDone(context.TODO(), 1, mockTask)

		assert.NoError(t, err)
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

func TestTaskDelete(t *testing.T) {
	mockTaskRepo := new(mocks.Repository)
	//nowTime := time.Now().UTC()

	mockTask := &models.Task{
		ID:      2,
		Name:    "task one",
		Status:  "pending",
		Comment: "nai",
		//UpdatedAt: &nowTime,
		//CreatedAt: &nowTime,
		UserID: 1,
	}

	t.Run("success", func(t *testing.T) {

		mockTaskRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).
			Return(mockTask, nil).Once()
		mockTaskRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).
			Return(nil).Once()

		u := NewTaskUsecase(mockTaskRepo, time.Second*2)
		err := u.Delete(context.TODO(), 1)

		assert.NoError(t, err)
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
