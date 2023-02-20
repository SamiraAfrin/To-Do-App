package usecase

import (

	"context"
	"time"
	
	"To_Do_App/models"
	"To_Do_App/Task"

)

type taskUsecase struct {
	taskRepo 		Task.Repository
	contextTimeout  time.Duration
}

func NewTaskUsecase(t Task.Repository, timeout time.Duration) Task.Usecase {

	return &taskUsecase {
		taskRepo: 		t,
		contextTimeout: timeout,
	}
}

// Delete the task using task id
func (t *taskUsecase) Delete(c context.Context, task_id int64) error {

	ctx, cancel:= context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	existedTask, err := t.taskRepo.GetByID(ctx, task_id)
	if err != nil {
		return err
	}

	if existedTask == nil {
		return models.ErrNotFound
	}

	return t.taskRepo.Delete(ctx, task_id)
}

// fetch the task using task id
func (t *taskUsecase) GetByID(c context.Context, task_id int64) (*models.Task, error){

	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	res, err := t.taskRepo.GetByID(ctx, task_id)
	if err != nil{
		return nil, err
	}

	return res, nil

}

// fetch all the data using the user id
func (t *taskUsecase) GetByUserID(c context.Context, user_id int64) ([]*models.Task, error){
	
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	res, err := t.taskRepo.GetByUserID(ctx, user_id)
	if err != nil{
		return nil, err
	}

	return res, nil

}

// fetch all the task in the task database 
func (t *taskUsecase) GetAllTask(c context.Context) ([]*models.Task, error){
	 
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()
	
	res, err := t.taskRepo.GetAllTask(ctx)
	if err != nil{
		return nil, err
	}

	return res, nil
}

// Store the data in the task database
func (t *taskUsecase) Store(c context.Context, task *models.Task) error{
	
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	tm := time.Now()
	task.CreatedAt = &tm
	task.UpdatedAt = task.CreatedAt

	err := t.taskRepo.Store(ctx, task)
	if err != nil{
		return err
	}

	return nil
}

// Update the existing task
func (t *taskUsecase) Update(c context.Context, task *models.Task) error{

	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	tm := time.Now()
	task.UpdatedAt = &tm
	return t.taskRepo.Update(ctx, task)

}

// Update the task status
func (t *taskUsecase) UpdateDone(c context.Context, task_id int64, task *Task.TaskPatchReq) error{
	 
	ctx, cancel := context.WithTimeout(c, t.contextTimeout)
	defer cancel()

	_, err := t.taskRepo.GetByID(ctx, task_id)
	if err != nil{
		return err
	}

	tm := time.Now()
	task.UpdatedAt = &tm
	taskerr := t.taskRepo.UpdateDone(ctx, task_id, &models.Task{
		ID	  : task_id,
		Status: task.Status,
		UpdatedAt: task.UpdatedAt,
	})
	if taskerr != nil{
		return taskerr
	}

	return nil
	
}
