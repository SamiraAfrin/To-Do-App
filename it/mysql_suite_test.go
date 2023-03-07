package it_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	connStr string
	err     error
)

type MysqlRepositoryTestSuite struct {
	gormDB *gorm.DB
	ctx    context.Context
	suite.Suite
}

func (m *MysqlRepositoryTestSuite) SetupSuite() {

	connStr = "root:123@tcp(localhost:3306)/recordings?charset=utf8&parseTime=True&loc=Local"
	var err error
	m.gormDB, err = gorm.Open(mysql.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Can't connect to mysql", err)
	}
}
func TestMysqlRepositoryTestSuite(t *testing.T) {
	suite.Run(t, &MysqlRepositoryTestSuite{})
}
func (m *MysqlRepositoryTestSuite) SetupTest() {
}

func (m *MysqlRepositoryTestSuite) TearDownTest() {
}
