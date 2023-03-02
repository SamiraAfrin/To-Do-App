package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"time"

	// Echo Framewrokss
	"github.com/labstack/echo"

	//Viper
	"github.com/spf13/viper"

	//Database
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	//Repository
	_taskRepo "To_Do_App/Task/repository"
	_userRepo "To_Do_App/User/repository"

	//Usecase
	_taskUsecase "To_Do_App/Task/usecase"
	_userUsecase "To_Do_App/User/usecase"

	// Delivery
	_taskHttpDelivery "To_Do_App/Task/delivery/http"
)

func init() {

	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", viper.GetString(`database.user`),
		viper.GetString(`database.pass`), viper.GetString(`database.net`), viper.GetString(`database.host`),
		viper.GetInt(`database.port`), viper.GetString(`database.name`))
	log.Print("from main")
	spew.Dump(dsn)
	// Get a database handle.
	var err error
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	// Echo Framewrok
	e := echo.New()

	// // Repository
	taskRepo := _taskRepo.NewMysqlTaskRepo(dbConn)
	userRepo := _userRepo.NewMysqlUserRepo(dbConn)

	// //Usecase
	timeoutContext := 10 * time.Second
	taskUsecase := _taskUsecase.NewTaskUsecase(taskRepo, timeoutContext)
	// fmt.Println(taskUsecase)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, timeoutContext)

	// //Delivery
	_taskHttpDelivery.NewTaskHandler(e, taskUsecase, userUsecase)

	e.Logger.Fatal(e.Start(viper.GetString("server.address")))

}
