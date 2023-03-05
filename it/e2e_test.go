package it_test

import (
	"To_Do_App/server"
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
type TaskMessage struct {
	ID        int64
	Name      string
	Status    string
	Comment   string
	UpdatedAt *time.Time
	CreatedAt *time.Time
	UserID    int64
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
	reqStr := `{"name":"Task1", "status": "Cholche gari atrabari", "comment":"kichu j bolar nai", "updated_at": "2023-02-20T09:31:34+06:00", "created_at": "2023-02-20T09:31:34+06:00", "user_id": 2}`
	req, err := http.NewRequest(echo.POST, fmt.Sprintf("http://localhost:%d/tasks", s.port), strings.NewReader(reqStr))
	s.NoError(err)
	spew.Dump(reqStr)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	client := http.Client{}
	response, err := client.Do(req)
	s.NoError(err)
	s.Equal(http.StatusCreated, response.StatusCode)

	byteBody, err := io.ReadAll(response.Body)
	s.NoError(err)

	var treq TaskMessage // request
	err = json.Unmarshal([]byte(reqStr), &treq)
	if err != nil {
		log.Fatal("err caught")
	}

	var tres TaskMessage // response
	err = json.Unmarshal(byteBody, &tres)
	if err != nil {
		log.Fatal("err caught")
	}
	s.Equal(tres.Name, treq.Name)
	s.Equal(tres.Comment, treq.Comment)
	s.Equal(tres.CreatedAt, treq.CreatedAt)
	s.Equal(tres.UserID, treq.UserID)
	response.Body.Close()
}
