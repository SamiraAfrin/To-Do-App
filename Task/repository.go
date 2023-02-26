package Task

import (
	"To_Do_App/models"
	"context"
)

type Repository interface {
	Delete(ctx context.Context, task_id int64) error
	GetByID(ctx context.Context, task_id int64) (*models.Task, error)
	GetByUserID(ctx context.Context, user_id int64) ([]*models.Task, error)
	GetAllTask(ctx context.Context) ([]*models.Task, error)
	Store(ctx context.Context, a *models.Task) error
	Update(ctx context.Context, a *models.Task) error
	UpdateDone(ctx context.Context, task_id int64, task_status *models.Task) error
}
