package repository

import (
	"context"

	"github.com/meles-z/golang-graphql/app/models"
)

type MovieRepository interface {
	Create(ctx context.Context, input models.MovieInput) (*models.Movie, error)
	Update(ctx context.Context, id string, input models.MovieInput) (*models.Movie, error)
	Delete(ctx context.Context, id string) (bool, error)
	GetByID(ctx context.Context, id string) (*models.Movie, error)
	GetAll(ctx context.Context, filter *models.MovieFilter) ([]*models.Movie, error)
}
