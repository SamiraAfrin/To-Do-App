package server

import (
	_taskHttpDelivery "To_Do_App/Task/delivery/http"
	taskRepo "To_Do_App/Task/repository"
	_taskUsecase "To_Do_App/Task/usecase"
	userRepo "To_Do_App/User/repository"
	_userUsecase "To_Do_App/User/usecase"
	"context"
	"fmt"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	Port       int
	DBConn     *gorm.DB
	SeverReady chan bool
}

func (s *Server) Start() {
	appPort := fmt.Sprintf(":%d", 8000)

	taskRepo := taskRepo.NewMysqlTaskRepo(s.DBConn)
	userRepo := userRepo.NewMysqlUserRepo(s.DBConn)

	// //Usecase
	timeoutContext := 10 * time.Second
	taskUsecase := _taskUsecase.NewTaskUsecase(taskRepo, timeoutContext)
	// fmt.Println(taskUsecase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)

	e := echo.New()
	_taskhandler, _ := _taskHttpDelivery.NewTaskHandler(e, taskUsecase, userUsecase)
	e.POST("/tasks", _taskhandler.Store)

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
