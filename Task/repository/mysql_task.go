package repository

import (

	"context"
	"gorm.io/gorm"
	"errors"
	
	"To_Do_App/Task"
	"To_Do_App/models"
	


)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlTaskRepo struct{
	Conn *gorm.DB
}



func NewMysqlTaskRepo(db *gorm.DB) Task.Repository{

	return &mysqlTaskRepo{
		Conn: db,
	}
}

// Delete the task for the specific task id
func (m *mysqlTaskRepo) Delete(ctx context.Context, task_id int64) error{

	var task models.Task
	result:= m.Conn.Delete(&task, task_id)

	return result.Error
	 
	
}

// Get a specific Task using task id 
func(m *mysqlTaskRepo) GetByID(ctx context.Context, task_id int64) (res *models.Task, err error){

	var task models.Task
	result:= m.Conn.First(&task, task_id)
	if result.Error != nil {
		return &task, errors.New("This task doesn't exist")
	}

	return &task, result.Error
	
}

func (m *mysqlTaskRepo) GetAllTask(ctx context.Context) ([]*models.Task, error) {

	var result []*models.Task
	query := m.Conn.Find(&result)
	if len(result) == 0  {
		return result, errors.New("Currently task table doesn't have any data")
	}
	
	return result, query.Error
	
}

// Fetch all the data using user_id 
 func (m *mysqlTaskRepo) GetByUserID(ctx context.Context, user_id int64) ([]*models.Task, error){

	var result []*models.Task
	query := m.Conn.Where("user_id = ?", user_id).Find(&result)
	if len(result) == 0  {
		return result, errors.New("This user don't have any task")
	}
	
	return result, query.Error
}


// Store new data in the database
 func (m *mysqlTaskRepo) Store(ctx context.Context, task *models.Task) error{

	result := m.Conn.Create(&task)

	return result.Error

	}	

// Update the existing task 
func (m *mysqlTaskRepo) Update(ctx context.Context, task *models.Task) error{

	result:= m.Conn.Save(&task)

	return result.Error
	

}

//Patch --> only status update 
func (m *mysqlTaskRepo) UpdateDone(ctx context.Context, task_id int64, task *models.Task) error {

	result := m.Conn.Model(&task).Select("Status", "UpdatedAt").Updates(models.Task{Status: task.Status, UpdatedAt: task.UpdatedAt})

	return result.Error

}






