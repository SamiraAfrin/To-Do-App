package it_test

import (
	"To_Do_App/Task"
	"To_Do_App/models"
	"To_Do_App/server"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

type e2eTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	dbConn          *gorm.DB
}

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}

func (s *e2eTestSuite) SetupSuite() {
	s.dbConnectionStr = "root:123@tcp(localhost:3306)/recordings?charset=utf8&parseTime=True&loc=Local"
	var err error
	s.dbConn, err = gorm.Open(mysql.Open(s.dbConnectionStr), &gorm.Config{})
	//s.Require().Error(err)
	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}

	serverReady := make(chan bool)

	server := server.Server{
		Port:       8000,
		DBConn:     s.dbConn,
		SeverReady: serverReady,
	}
	go server.Start()
	<-serverReady
}

func (s *e2eTestSuite) TearDownSuite() {
}
func (s *e2eTestSuite) SetupTest() {
	s.port = 8000
}
func (s *e2eTestSuite) TearDownTest() {
}

func (s *e2eTestSuite) Test_EndToEnd_CreateTask() {
	reqStr := `{"name":"Task1", "status": "Cholche gari jatrabari", "comment":"kichu j bolar nai", "updated_at": "2023-02-20T09:31:34+06:00", "created_at": "2023-02-20T09:31:34+06:00", "user_id": 2}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/tasks", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)

	var treq models.Task // request
	err = json.Unmarshal([]byte(reqStr), &treq)
	if err != nil {
		log.Fatal("err caught")
	}

	var tres models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}
	s.Equal(tres.Name, treq.Name)
	s.Equal(tres.Comment, treq.Comment)
	s.Equal(tres.Status, treq.Status)
	s.Equal(tres.UserID, treq.UserID)
	response.Body.Close()
}
func (s *e2eTestSuite) Test_EndToEnd_GetTaskByTaskID() {
	nowTime := time.Now().UTC()
	task := models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	s.NoError(s.dbConn.Create(&task).Error)
	tasks := GetAllTask()
	id := tasks[0].ID
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/tasks/%d", s.port, id), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)
	var tres models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}

	s.Equal(tres.ID, id)
	s.Equal(tres.Name, task.Name)
	s.Equal(tres.Comment, task.Comment)
	s.Equal(tres.Status, task.Status)
	s.Equal(tres.UserID, task.UserID)

	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllTask() {
	nowTime := time.Now().UTC()
	task := models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	s.NoError(s.dbConn.Create(&task).Error)
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/tasks", s.port), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)
	var tres []models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}

	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetTasksByUserID() {
	nowTime := time.Now().UTC()
	task := models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	s.NoError(s.dbConn.Create(&task).Error)
	tasks := GetAllTask()
	t := tasks[0]

	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/tasks/user/%d", s.port, t.UserID), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)
	var tres []models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}

	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_DeleteTaskByID() {
	nowTime := time.Now().UTC()
	task := models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	s.NoError(s.dbConn.Create(&task).Error)
	tasks := GetAllTask()

	// will delete last inserted data
	t := tasks[0]

	req, err := http.NewRequest(echo.DELETE, fmt.Sprintf("http://localhost:%d/tasks/%d", s.port, t.ID), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)
	var tres models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}
	s.Equal(http.StatusOK, response.StatusCode)

	//For checking
	// tasks = GetAllTask()
	//t = tasks[0]
	//spew.Dump(t.ID)
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_UpdateTaskByID() {
	nowTime := time.Now().UTC()
	task := models.Task{
		Name:      "First Taskss",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	updatedTask := models.Task{
		Name:      "Updated Task",
		Status:    "in progresss",
		Comment:   "kich nais",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    2,
	}
	s.NoError(s.dbConn.Create(&task).Error)
	tasks := GetAllTask()

	// will update last inserted data
	t := tasks[len(tasks)-1]
	updatedTask.ID = t.ID
	b, err := json.Marshal(&updatedTask)
	s.NoError(err)

	req, err := http.NewRequest(echo.PUT, fmt.Sprintf("http://localhost:%d/tasks/%d", s.port, t.ID), bytes.NewReader(b))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	//spew.Dump(string(byteBody))
	s.NoError(err)

	var tres models.Task // response

	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		spew.Dump(err)
		log.Fatal("err caught hocche")
	}
	s.Equal(http.StatusOK, response.StatusCode)

	//For checking
	//tasks = GetAllTask()
	//t = tasks[0]
	//spew.Dump(t.ID)
	//spew.Dump(t.Name)
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_CompleteUpdateTaskByID() {
	nowTime := time.Now().UTC()
	task := models.Task{
		Name:      "First Taskss",
		Status:    "in progress cholche",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    2,
	}
	updatedTask := Task.TaskPatchReq{
		Status:    "Done",
		UpdatedAt: &nowTime,
	}
	s.NoError(s.dbConn.Create(&task).Error)
	tasks := GetAllTask()

	// will update last inserted data
	t := tasks[0]
	b, err := json.Marshal(&updatedTask)
	s.NoError(err)
	//testing
	spew.Dump(t.ID)
	spew.Dump(t.Status)

	req, err := http.NewRequest(echo.PATCH, fmt.Sprintf("http://localhost:%d/tasks/%d", s.port, t.ID), bytes.NewReader(b))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	//spew.Dump(string(byteBody))
	s.NoError(err)

	var tres models.Task // response

	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		spew.Dump(err)
		log.Fatal("err caught hocche")
	}
	s.Equal(http.StatusOK, response.StatusCode)

	//For checking
	//tasks = GetAllTask()
	//t = tasks[0]
	//spew.Dump(t.ID)
	//spew.Dump(t.Status)
	response.Body.Close()
}

func GetAllTask() ([]models.Task) {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/tasks", 8000), nil)

	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if 200 != response.StatusCode {
		log.Fatal("http status code doesn't match")
	}

	byteBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var tres []models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}
	response.Body.Close()
	return tres
}
func (s *e2eTestSuite) Test_EndToEnd_CreateUser() {
	reqStr := `{"name":"User1"}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/users", s.port), strings.NewReader(reqStr))
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)

	var treq models.User // request
	err = json.Unmarshal([]byte(reqStr), &treq)
	if err != nil {
		log.Fatal("err caught")
	}

	var tres models.User // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}
	s.Equal(tres.Name, treq.Name)
	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_GetAllUser() {
	user := models.User{
		Name: "First User",
	}
	s.NoError(s.dbConn.Create(&user).Error)
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/users", s.port), nil)
	s.NoError(err)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)
	var tres []models.Task // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}

	response.Body.Close()
}

func (s *e2eTestSuite) Test_EndToEnd_UpdateUserByUserID() {
	user := models.User{
		Name: "First User",
	}
	updatedUser := models.User{
		Name: "Updated User",
	}
	s.NoError(s.dbConn.Create(&user).Error)
	users := GetAllUser()

	// will update last inserted data
	u := users[0]
	updatedUser.ID = u.ID
	b, err := json.Marshal(&updatedUser)
	s.NoError(err)
	//testing
	spew.Dump(u.ID)
	spew.Dump(u.Name)
	//
	req, err := http.NewRequest(echo.PUT, fmt.Sprintf("http://localhost:%d/users/%d", s.port, u.ID), bytes.NewReader(b))
	s.NoError(err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	//spew.Dump(string(byteBody))
	s.NoError(err)

	var tres models.User // response

	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		spew.Dump(err)
		log.Fatal("err caught hocche")
	}
	s.Equal(http.StatusOK, response.StatusCode)

	//For checking
	users = GetAllUser()
	u = users[0]
	spew.Dump(u.ID)
	spew.Dump(u.Name)
	response.Body.Close()
}

func GetAllUser() []models.User {
	req, err := http.NewRequest(echo.GET, fmt.Sprintf("http://localhost:%d/users", 8000), nil)

	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if 200 != response.StatusCode {
		log.Fatal("http status code doesn't match")
	}

	byteBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var tres []models.User // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}
	response.Body.Close()
	return tres
}

