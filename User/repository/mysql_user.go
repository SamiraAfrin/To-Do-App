package repository

import (

	"context"
	// "database/sql"
	// "fmt"
	// "github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"To_Do_App/User"
	"To_Do_App/models"
	
)

type mysqlUserRepo struct {
	Conn *gorm.DB
}

func NewMysqlUserRepo(db *gorm.DB) User.Repository{

	return &mysqlUserRepo{
		Conn: db,
	}
}

// Store new data in the database
func (m *mysqlUserRepo) Store(ctx context.Context, user *models.User) error{

	result := m.Conn.Create(&user)
	
	return result.Error

}	

// Update the existing user
func (m *mysqlUserRepo) Update(ctx context.Context, user *models.User) error{

	result:= m.Conn.Save(&user)

	return result.Error
	
	
}

// Get all the user in the database
func (m *mysqlUserRepo) GetAllUser(ctx context.Context) ([]*models.User, error) {

	// var result []*models.TaskDB
	// d:= m.Conn.Find(&result)
	// fmt.Println(result)

	result := make([]*models.User, 0)
	d:= m.Conn.Find(&result)

	return result, d.Error

	
}




