package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"

	"To_Do_App/Task"
	"To_Do_App/User"
	"To_Do_App/models"
)

type ResponseError struct {
	Message string `json:"message"`
}

type TaskHandler struct {
	TaskUsecase Task.Usecase
}

type UserHandler struct {
	UserUsecase User.Usecase
}

func NewTaskHandler(e *echo.Echo, task Task.Usecase, user User.Usecase) (*TaskHandler, *UserHandler) {

	taskhandler := &TaskHandler{
		TaskUsecase: task,
	}
	userhandler := &UserHandler{
		UserUsecase: user,
	}

	e.DELETE("/tasks/:ID", taskhandler.Delete)
	e.POST("/tasks", taskhandler.Store)
	e.GET("/tasks/:ID", taskhandler.GetByID)
	e.GET("/tasks/user/:userID", taskhandler.GetByUserID)
	e.GET("tasks", taskhandler.GetAllTask)
	e.PUT("tasks/:ID", taskhandler.Update)
	e.PATCH("tasks/:ID", taskhandler.UpdateDone)

	e.POST("/users", userhandler.UserStore)
	e.PUT("users/:ID", userhandler.UserUpdate)
	e.GET("users", userhandler.GetAllUser)

	return taskhandler, userhandler
}

func (t *TaskHandler) Delete(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = t.TaskUsecase.Delete(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: "Delete Successfully"})
}

func (t *TaskHandler) Store(c echo.Context) error {

	var task models.Task
	err := c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&task); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = t.TaskUsecase.Store(ctx, &task)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, task)

}

func (t *TaskHandler) GetByID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tsk, err := t.TaskUsecase.GetByID(ctx, id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, tsk)
}

func (t *TaskHandler) GetByUserID(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	user_id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	tsk, err := t.TaskUsecase.GetByUserID(ctx, user_id)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, tsk)
}

func (t *TaskHandler) GetAllTask(c echo.Context) error {

	ctx := c.Request().Context()
	listAr, err := t.TaskUsecase.GetAllTask(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listAr)
}

func (t *TaskHandler) Update(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	var task models.Task
	task.ID = id
	err = c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isRequestValid(&task); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = t.TaskUsecase.Update(ctx, &task)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: "Updated"})

}

func (t *TaskHandler) UpdateDone(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		//log.Print(err)
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var req *Task.TaskPatchReq
	err = c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrNotFound.Error())
	}

	err = t.TaskUsecase.UpdateDone(ctx, id, req)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: "Update done for status"})

}

func (u *UserHandler) UserStore(c echo.Context) error {

	var user models.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isUserRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = u.UserUsecase.Store(ctx, &user)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, user)

}

func (u *UserHandler) UserUpdate(c echo.Context) error {

	idP, err := strconv.Atoi(c.Param("ID"))
	if err != nil {
		return c.JSON(http.StatusNotFound, models.ErrNotFound.Error())
	}

	id := int64(idP)
	var user models.User
	user.ID = id
	err = c.Bind(&user)
	if err != nil {
		//log.Print("HELLO")
		//log.Print(err)
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if ok, err := isUserRequestValid(&user); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	err = u.UserUsecase.Update(ctx, &user)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, ResponseError{Message: "Updated"})

}

func (u *UserHandler) GetAllUser(c echo.Context) error {

	ctx := c.Request().Context()
	listAr, err := u.UserUsecase.GetAllUser(ctx)
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, listAr)
}

// Utility Functions
func isRequestValid(task *models.Task) (bool, error) {

	validate := validator.New()
	err := validate.Struct(task)
	if err != nil {
		return false, err
	}

	return true, nil
}

func isUserRequestValid(user *models.User) (bool, error) {

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		return false, err
	}

	return true, nil
}

func getStatusCode(err error) int {

	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
