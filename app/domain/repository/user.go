package repository

import "github.com/meles-z/golang-graphql/app/models"

type UserRepository interface {
	CreateUser(*models.User) (*models.User, error)
	GetUsers() []*models.User
	GetUserById(id string) (*models.User, error)
	UpdateUser(*models.User) (*models.User, error)
	DeleteUser(id string) error
}
