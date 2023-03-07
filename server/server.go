package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
	
	_taskHttpDelivery "To_Do_App/Task/delivery/http"
	taskRepo "To_Do_App/Task/repository"
	_taskUsecase "To_Do_App/Task/usecase"
	userRepo "To_Do_App/User/repository"
	_userUsecase "To_Do_App/User/usecase"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Server struct {
	Port       int
	DBConn     *gorm.DB
	SeverReady chan bool
}

func (s *Server) Start() {

	appPort := fmt.Sprintf(":%d", 8000)

	//Repository
	taskRepo := taskRepo.NewMysqlTaskRepo(s.DBConn)
	userRepo := userRepo.NewMysqlUserRepo(s.DBConn)

	//Usecase
	timeoutContext := 10 * time.Second
	taskUsecase := _taskUsecase.NewTaskUsecase(taskRepo, timeoutContext)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)

	e := echo.New()

	// Delivery
	_taskhandler, _userhandler := _taskHttpDelivery.NewTaskHandler(e, taskUsecase, userUsecase)

	// For Task
	e.POST("/tasks", _taskhandler.Store)
	e.GET("/tasks/:ID", _taskhandler.GetByID)
	e.GET("tasks", _taskhandler.GetAllTask)
	e.GET("/tasks/user/:userID", _taskhandler.GetByUserID)
	e.DELETE("/tasks/:ID", _taskhandler.Delete)
	e.PUT("tasks/:ID", _taskhandler.Update)

	// For User
	e.POST("/users", _userhandler.UserStore)
	e.PUT("users/:ID", _userhandler.UserUpdate)
	e.GET("users", _userhandler.GetAllUser)

	go func() {
		if err := e.Start(appPort); err != nil {
			logrus.Errorf(err.Error())
			logrus.Infof("shutting down the server")
		}
	}()

	if s.SeverReady != nil {
		s.SeverReady <- true
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		logrus.Fatalf("failed to gracefully shutdown the server: %s", err)
	}
}
