package it_test

import (
	"time"

	"To_Do_App/Task/repository"
	"To_Do_App/models"
)

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_Store() {

	nowTime := time.Now().UTC()
	newTask := &models.Task{
		Name:      "Ami",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    2,
	}

	repo := repository.NewMysqlTaskRepo(m.gormDB)
	m.Assert().NoError(repo.Store(m.ctx, newTask))

	id := newTask.ID
	var result models.Task

	m.Assert().NoError(m.gormDB.First(&result, models.Task{ID: id}).Error)
	m.Assert().Equal(newTask.ID, result.ID)
	m.Assert().Equal(newTask.Name, result.Name)
	m.Assert().Equal(newTask.Status, result.Status)
	m.Assert().Equal(newTask.Comment, result.Comment)
	m.Assert().Equal(newTask.UserID, result.UserID)
}

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_GetByID() {

	searchTask := &models.Task{
		Name:    "Ami",
		Status:  "in progress",
		Comment: "kich nai",
		UserID:  2,
	}

	r := repository.NewMysqlTaskRepo(m.gormDB)

	m.Assert().NoError(r.Store(m.ctx, searchTask))
	result, err := r.GetByID(m.ctx, searchTask.ID)

	m.Assert().NoError(err)
	m.Assert().Equal(searchTask.ID, result.ID)
	m.Assert().Equal(searchTask.UserID, result.UserID)

}

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_GetAllTask() {

	nowTime := time.Now().UTC()
	firstTask := &models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    2,
	}

	secondTask := &models.Task{
		Name:      "Second Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}

	r := repository.NewMysqlTaskRepo(m.gormDB)

	m.Assert().NoError(r.Store(m.ctx, firstTask))
	m.Assert().NoError(r.Store(m.ctx, secondTask))

	_, err := r.GetAllTask(m.ctx)
	m.Assert().NoError(err)

}

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_GetByUserID() {

	nowTime := time.Now().UTC()
	firstTask := &models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}

	r := repository.NewMysqlTaskRepo(m.gormDB)
	m.Assert().NoError(r.Store(m.ctx, firstTask))

	_, err := r.GetByUserID(m.ctx, firstTask.UserID)
	m.Assert().NoError(err)

}

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_Delete() {

	nowTime := time.Now().UTC()
	Task := &models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}

	r := repository.NewMysqlTaskRepo(m.gormDB)
	m.Assert().NoError(r.Store(m.ctx, Task))

	err := r.Delete(m.ctx, Task.ID)
	m.Assert().NoError(err)

}

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_Update() {
	nowTime := time.Now().UTC()
	Task := &models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	UpdatedTask := &models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    3,
	}

	r := repository.NewMysqlTaskRepo(m.gormDB)
	m.Assert().NoError(r.Store(m.ctx, Task))

	err := r.Update(m.ctx, UpdatedTask)
	m.Assert().NoError(err)

}

func (m *MysqlRepositoryTestSuite) TestMysqlTaskRepository_UpdateDone() {

	nowTime := time.Now().UTC()
	Task := &models.Task{
		Name:      "First Task",
		Status:    "in progress",
		Comment:   "kich nai",
		UpdatedAt: &nowTime,
		CreatedAt: &nowTime,
		UserID:    1,
	}
	UpdatedTask := &models.Task{
		Status:    "in progress",
		UpdatedAt: &nowTime,
	}

	r := repository.NewMysqlTaskRepo(m.gormDB)
	m.Assert().NoError(r.Store(m.ctx, Task))

	UpdatedTask.ID = Task.ID

	err := r.UpdateDone(m.ctx, Task.ID, UpdatedTask)
	m.Assert().NoError(err)

}
