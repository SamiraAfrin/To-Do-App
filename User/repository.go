package User

import (
	"To_Do_App/models"
	"context"
)

type Repository interface {
	Store(ctx context.Context, a *models.User) error
	Update(ctx context.Context, a *models.User) error
	GetAllUser(ctx context.Context) ([]*models.User, error)
}
