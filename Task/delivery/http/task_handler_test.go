package http

import (
	"To_Do_App/Task/mocks"
	"To_Do_App/models"
	"encoding/json"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

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
	c.SetPath("tasks/:id")
	c.SetParamNames("id")
	c.SetParamValues(strconv.Itoa(num))
	handler := TaskHandler{
		TaskUsecase: mockUCase,
	}
	err = handler.Delete(c)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}
