package models

// User represent the user model
type User struct {
	ID   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
