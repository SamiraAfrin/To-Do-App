package User

import (
	
	"context"
	"To_Do_App/models"

)

type Repository interface {

	Store(ctx context.Context, a *models.User) error
	Update(ctx context.Context, a *models.User) error
	GetAllUser(ctx context.Context) ([]*models.User, error)

}