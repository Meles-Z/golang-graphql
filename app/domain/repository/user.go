package repository

import (
	"context"

	"github.com/meles-z/golang-graphql/app/models"
)

type UserRepository interface {
	Create(ctx context.Context, input models.UserInput) (*models.User, error)
	Update(ctx context.Context, id string, input models.UserInput) (*models.User, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context, filter *models.UserFilter) ([]*models.User, error)
}
