package http

import (
	"To_Do_App/Task/mocks"
	userMock "To_Do_App/User/mocks"
	"To_Do_App/models"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// for the user
func TestCreateUser(t *testing.T) {
	mockUser := models.User{
		Name: "Hello",
	}

	tempMockUser := mockUser
	tempMockUser.ID = 15

	j, err := json.Marshal(tempMockUser)
	assert.NoError(t, err)

	mockUCase := new(userMock.Usecase)
	//spew.Dump(mock.Anything, mock.AnythingOfType("*models.User"), "hello")
	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/users", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users")

	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.UserStore(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestGetAllUser(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	mockUCase := new(userMock.Usecase)
	mockListUser := make([]*models.User, 0)
	mockListUser = append(mockListUser, &mockUser)

	mockUCase.On("GetAllUser", mock.Anything).Return(mockListUser, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "users", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.GetAllUser(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestUserUpdate(t *testing.T) {
	var mockUser models.User
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)

	j, err := json.Marshal(mockUser)
	assert.NoError(t, err)
	spew.Dump(j)

	mockUCase := new(userMock.Usecase)

	num := int(mockUser.ID)
	spew.Dump(num)
	spew.Dump(mockUser)
	//spew.Dump(mock.Anything, mock.AnythingOfType("*models.User"))
	mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*models.User")).Return(nil)

	e := echo.New()

	req, err := http.NewRequest(echo.PUT, "/users/"+strconv.Itoa(num), strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("users/:ID")
	c.SetParamNames("ID")
	c.SetParamValues(strconv.Itoa(num))
	handler := UserHandler{
		UserUsecase: mockUCase,
	}
	err = handler.UserUpdate(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

// for the task

func TestCreateTask(t *testing.T) {

	nowTime := time.Now().UTC()
	mockTask := models.Task{
		Name:      "Hello",
		Status:    "progress",
		Comment:   "kisu nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    2,
	}

	tempMockTask := mockTask
	tempMockTask.ID = 15

	j, err := json.Marshal(tempMockTask)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	mockUCase.On("Store", mock.Anything, mock.AnythingOfType("*models.Task")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/tasks", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/tasks")

	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.Store(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestGetTaskByID(t *testing.T) {
	var mockTask models.Task
	err := faker.FakeData(&mockTask)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)

	num := int(mockTask.ID)
	mockUCase.On("GetByID", mock.Anything, int64(num)).Return(&mockTask, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tasks/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tasks/:ID")
	c.SetParamNames("ID")
	c.SetParamValues(strconv.Itoa(num))
	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.GetByID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetAllTask(t *testing.T) {
	var mockTask models.Task
	err := faker.FakeData(&mockTask)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	mockListTask := make([]*models.Task, 0)
	mockListTask = append(mockListTask, &mockTask)

	mockUCase.On("GetAllTask", mock.Anything).Return(mockListTask, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "tasks", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.GetAllTask(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestDelete(t *testing.T) {
	var mockTask models.Task
	err := faker.FakeData(&mockTask)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)

	num := int(mockTask.ID)

	//mockUCase.On("GetByID", mock.Anything, int64(num)).Return(nil)
	mockUCase.On("Delete", mock.Anything, int64(num)).Return(nil)

	e := echo.New()

	req, err := http.NewRequest(echo.DELETE, "/tasks/"+strconv.Itoa(num), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tasks/:ID")
	c.SetParamNames("ID")
	c.SetParamValues(strconv.Itoa(num))
	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestGetTasksByUserID(t *testing.T) {
	var mockTask models.Task
	err := faker.FakeData(&mockTask)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	mockListTask := make([]*models.Task, 0)

	mockListTask = append(mockListTask, &mockTask)
	userID := int(mockTask.UserID)
	mockUCase.On("GetByUserID", mock.Anything, int64(userID)).Return(mockListTask, nil)

	e := echo.New()
	req, err := http.NewRequest(echo.GET, "/tasks/user/"+strconv.Itoa(userID), strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("tasks/user/:userID")
	c.SetParamNames("userID")
	c.SetParamValues(strconv.Itoa(userID))

	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.GetByUserID(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestUpdate(t *testing.T) {
	nowTime := time.Now().UTC()
	mockTask := models.Task{
		Name:      "Hello",
		Status:    "progress",
		Comment:   "kisu nai",
		UpdatedAt: &nowTime,
		UserID:    2,
	}
	num := int(mockTask.ID)
	j, err := json.Marshal(mockTask)
	assert.NoError(t, err)

	mockUCase := new(mocks.Usecase)
	mockUCase.On("Update", mock.Anything, mock.AnythingOfType("*models.Task")).Return(nil)

	e := echo.New()
	req, err := http.NewRequest(echo.PUT, "/tasks/"+strconv.Itoa(num), strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("tasks/:ID")
	c.SetParamNames("ID")
	c.SetParamValues(strconv.Itoa(num))

	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.Update(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)

}

//func TestCompleteTask(t *testing.T) {
//	nowTime := time.Now().UTC()
//
//	mockTask := models.Task{}
//	modelTask := &Task.TaskPatchReq{
//		Status:    "pending",
//		UpdatedAt: &nowTime,
//	}
//	mockTask.Status = modelTask.Status
//	mockTask.UpdatedAt = modelTask.UpdatedAt
//	num := int(mockTask.ID)
//	j, err := json.Marshal(mockTask)
//	assert.NoError(t, err)
//
//	mockUCase := new(mocks.Usecase)
//	mockUCase.On("UpdateDone", mock.Anything, mock.AnythingOfType("int64"), mock.AnythingOfType("*models.Task")).Return(nil)
//
//	e := echo.New()
//	req, err := http.NewRequest(echo.PATCH, "/tasks/"+strconv.Itoa(num), strings.NewReader(string(j)))
//	assert.NoError(t, err)
//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//	rec := httptest.NewRecorder()
//	c := e.NewContext(req, rec)
//	c.SetPath("tasks/:ID")
//	c.SetParamNames("ID")
//	c.SetParamValues(strconv.Itoa(num))
//
//	handler := TaskHandler{
//		TaskUsecase: mockUCase,
//	}
//	err = handler.UpdateDone(c)
//	require.NoError(t, err)
//
//	assert.Equal(t, http.StatusOK, rec.Code)
//	mockUCase.AssertExpectations(t)
//
//}
