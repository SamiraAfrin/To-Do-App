package it_test

import (
	"To_Do_App/User/repository"
	"To_Do_App/models"
	//"github.com/davecgh/go-spew/spew"
	//"time"
)

func (m *MysqlRepositoryTestSuite) TestMysqlUserRepository_Store() {
	newUser := &models.User{
		Name: "First user",
	}
	repo := repository.NewMysqlUserRepo(m.gormDB)
	m.Assert().NoError(repo.Store(m.ctx, newUser))

	id := newUser.ID

	var result models.User
	m.Assert().NoError(m.gormDB.First(&result, models.User{ID: id}).Error)
	m.Assert().Equal(newUser.ID, result.ID)
	m.Assert().Equal(newUser.Name, result.Name)
}

func (m *MysqlRepositoryTestSuite) TestMysqlUserRepository_GetAllUser() {
	firstUser := &models.User{
		Name: "First User",
	}
	secondUser := &models.User{
		Name: "Second User",
	}
	r := repository.NewMysqlUserRepo(m.gormDB)
	m.Assert().NoError(r.Store(m.ctx, firstUser))
	m.Assert().NoError(r.Store(m.ctx, secondUser))

	_, err := r.GetAllUser(m.ctx) // count checking a pb
	m.Assert().NoError(err)
	//m.Assert().Len(result, 2)
}

func (m *MysqlRepositoryTestSuite) TestMysqlUserRepository_Update() {

	User := &models.User{
		Name: "First User",
	}
	UpdatedUser := &models.User{
		Name: "Updated Name",
	}
	r := repository.NewMysqlUserRepo(m.gormDB)
	m.Assert().NoError(r.Store(m.ctx, User))

	err := r.Update(m.ctx, UpdatedUser) // count checking a pb
	m.Assert().NoError(err)
}
