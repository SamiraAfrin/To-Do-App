package User

import (

	"context"
	"To_Do_App/models"
	
)




type Usecase interface {

	Store(ctx context.Context, a *models.User) error
	Update(ctx context.Context, a *models.User) error	
	GetAllUser(c context.Context) ([]*models.User, error)

}