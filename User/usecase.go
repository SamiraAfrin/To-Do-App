package User

import (
	"To_Do_App/models"
	"context"
)

type Usecase interface {
	Store(ctx context.Context, a *models.User) error
	Update(ctx context.Context, a *models.User) error
	GetAllUser(c context.Context) ([]*models.User, error)
}
