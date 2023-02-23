package mocks

import (
	context "context"
	mock "github.com/stretchr/testify/mock"
	models "To_Do_App/models"
	"To_Do_App/Task"


)

type Usecase struct {
	mock.Mock
}


// Delete provides a mock function with given fields: ctx, id
func (_m *Usecase) Delete (ctx context.Context, task_id int64) error {

	ret := _m.Called(ctx,task_id)
	
	var err error
	if rf, ok :=  ret.Get(0).func(context.Context, int64) error); ok {
		err = rf(ctx, id)
	}else{
		err = ret.Error(0)
	}
	return err
}


// Fetch provides a mock function with given fields: ctx, task_id
func (_m *Usecase) GetByID(ctx context.Context, task_id int64)(*models.Task, error){

	ret := _m.Called(ctx, task_id)

	var res *models.Task
	if rf, ok := ret.Get(0).func(context.Context, int64); ok{
		res = rf(ctx,task_id)
	}else{
		if ret.Get(0) != nil{
			res = ret.Get(0).(*models.Task)
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok{
		err = rf(ctx,task_id)
	}else{
		err = ret.Error(1)
	}

	return res, err
}


// Fetch provides a mock function with given fields: ctx, user_id
func (_m *Usecase) GetByUserID(ctx context.Context, user_id int64) ([]*models.Task, error){

	ret := _m.Called(ctx, user_id)

	var res []*models.Task
	if rf, ok := ret.Get(0).func(context.Context, int64); ok{
		res = rf(ctx,user_id)
	}else{
		if ret.Get(0) != nil{
			res = ret.Get(0).(*models.Task)
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok{
		err = rf(ctx,user_id)
	}else{
		err = ret.Error(1)
	}

	return res, err
}

func (_m *Usecase) GetAllTask(ctx context.Context) ([]*models.Task, error){

	ret := _m.Called(ctx)

	var res []*models.Task
	if rf, ok := ret.Get(0).func(context.Context, int64); ok{
		res = rf(ctx)
	}else{
		if ret.Get(0) != nil{
			res = ret.Get(0).(*models.Task)
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok{
		err = rf(ctx)
	}else{
		err = ret.Error(1)
	}

	return res, err
}

func (_m *Usecase) Store(ctx context.Context, task *models.Task) error{

	ret := _m.Called(ctx, task)

	var err error 
	if rf, ok := ret.Get(0).(func(context.Context,*models.Task) error); ok{
		err = rf(ctx, task)
	} else{
		err = ret.Error(0)
	}

	return err
}

func (_m *Usecase) Update(ctx context.Context, task *models.Task) error{
	
	ret := _m.Called(ctx, task)
	
	var err error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Task)error); ok{
		err = rf(ctx, task)
	}else{
		err = ret.Error(0)
	}

	return err 
}

func (_m *Usecase) Update(ctx context.Context, task_id int64, statusReq *Task.TaskPatchReq) error{
	
	ret := _m.Called(ctx, task_id, statusReq)
	
	var err error
	if rf, ok := ret.Get(0).(func(context.Context, int64, *Task.TaskPatchReq)error); ok{
		err = rf(ctx, task_id, statusReq)
	}else{
		err = ret.Error(0)
	}

	return err 
}



